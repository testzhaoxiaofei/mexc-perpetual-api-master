package mexcperpetualapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	fhttp "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/models"
	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/services"
	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/utils"
)

type MexcPerpetualClient struct {
	uid    string // Cookie中的uid字段
	client tls_client.HttpClient

	mToken string // 基于uid生成的md5
	mHash  string // mToken的md5
	// TODO: 给dolosConfig加上过期检查
	dolosConfig  *models.DolosConfig  // 从接口中获取, 下单前要检查是否过期(过期时间为接口请求的秒数+86400秒)
	customerInfo *models.CustomerInfo // 从接口中获取
	// triggerProtect string               // 从 https://futures.mexc.com/api/v1/private/user_pref/get 获取的字段,放到下单之前获取这个字段
}

// 创建一个新的MexcPerpetualClient
//
// 由于初始化的信息走多个接口，且接口走多了后会影响数据的返回，不要频繁初始化
func NewMexcPerpetualClient(cookieUid, proxyUrl string) (*MexcPerpetualClient, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(10),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		tls_client.WithClientProfile(profiles.Chrome_131),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}

	if proxyUrl != "" {
		client.SetProxy(proxyUrl)
	}

	// 设置mToken
	mToken := utils.CalculateMD5(cookieUid)
	mhash := utils.CalculateMD5(mToken)

	mexcPerpetualClient := &MexcPerpetualClient{
		uid:    cookieUid,
		client: client,
		mToken: mToken,
		mHash:  mhash,
	}

	// 获取dolos配置
	if err := mexcPerpetualClient.getDolosConfig(); err != nil {
		return nil, err
	}

	// 获取用户信息
	if err := mexcPerpetualClient.getCustomerInfo(); err != nil {
		return nil, err
	}

	return mexcPerpetualClient, nil
}

// 开空
//
// @param symbol 合约名称, 如ETH_USDT
//
// @param vol 数量
//
// @return 订单号
func (mpc *MexcPerpetualClient) OpenKong(symbol string, vol float64, leverage int) (string, error) {
	OPEN_KONG_SIDE := 3

	return mpc.submitPerpetualOrder(symbol, vol, leverage, OPEN_KONG_SIDE)
}

// 开多
//
// @param symbol 合约名称, 如ETH_USDT
//
// @param vol 数量
//
// @return 订单号
func (mpc *MexcPerpetualClient) OpenDuo(symbol string, vol float64, leverage int) (string, error) {
	OPEN_DUO_SIDE := 1

	return mpc.submitPerpetualOrder(symbol, vol, leverage, OPEN_DUO_SIDE)
}

func (mpc *MexcPerpetualClient) submitPerpetualOrder(symbol string, vol float64, leverage, side int) (string, error) {
	// https://futures.mexc.com/api/v1/private/order/create?mhash=b2d52c26ce0879d5d704893fa6e026a6
	if mpc.dolosConfig == nil || mpc.customerInfo == nil {
		return "", errors.New("dolosConfig or customerInfo is nil")
	}
	orderSig, err := services.GenerateOrderCreateSig(mpc.mToken, mpc.customerInfo.MemberID, mpc.dolosConfig.Parameters)
	if err != nil {
		fmt.Println("OpenKong GenerateOrderCreateSig err:", err)
		return "", err
	}
	// 获取triggerProtect字段
	triggerProtect, err := mpc.getTriggerProtect()
	if err != nil {
		fmt.Println("OpenKong getTriggerProtect err:", err)
		return "", err
	}

	orderCreateReq := models.MexcOrderCreateReqeust{
		Symbol:        symbol,
		Vol:           vol,
		Side:          side,
		OpenType:      2,
		Type:          "5",
		Leverage:      leverage,
		MarketCeiling: false,
		PriceProtect:  triggerProtect,
		Mtoken:        mpc.mToken,
		Mhash:         mpc.mHash,
		Chash:         mpc.dolosConfig.Chash,
	}

	if orderSig == nil {
		return "", errors.New("orderSig is nil")
	}
	orderCreateReq.P0 = orderSig.P0
	orderCreateReq.K0 = orderSig.K0
	orderCreateReq.Ts = orderSig.Ts

	// orderBytes, err := json.Marshal(orderCreateReq)
	// if err == nil {
	// 	slog.Info("OpenKong", slog.String("orderCreateReq", string(orderBytes)))
	// }

	reqMapData := orderCreateReq.ToMapData()
	orderCreateUrl := "https://futures.mexc.com/api/v1/private/order/create?mhash=" + mpc.mHash
	cookies := map[string]string{
		"u_id": mpc.uid,
	}
	xMxcNonce, xMxcSign, err := services.HeaderSign(mpc.uid, reqMapData)
	if err != nil {
		fmt.Println("OpenKong HeaderSign err:", err)
		return "", err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Referer":       "https://futures.mexc.com/zh-MY/exchange/AIXBT_USDT?_from=search",
		"Authorization": mpc.uid,
		"x-mxc-nonce":   xMxcNonce,
		"x-mxc-sign":    xMxcSign,
		"User-Agent":    utils.USER_AGENT,
	}
	respCode, respBody, err := mpc.post(orderCreateUrl, headers, cookies, reqMapData)
	if err != nil {
		fmt.Println("OpenKong post err:", err)
		return "", err
	}
	defer respBody.Close()
	bodyBytes, err := io.ReadAll(respBody)
	if err != nil {
		fmt.Println("OpenKong err:", err)
		return "", err
	}

	if respCode != http.StatusOK {
		return "", errors.New("OpenKong respCode is " + strconv.Itoa(respCode))
	}

	// fmt.Println("OpenKong respBody:", string(bodyBytes))
	respStruct := models.MexcOrderCreateResponse{}
	if err := json.Unmarshal(bodyBytes, &respStruct); err != nil {
		fmt.Println("OpenKong unmarshall respStruct  err:", err)
		return "", err
	}

	if respStruct.Code != 0 {
		return "", errors.New("OpenKong respStruct.Code is " + strconv.Itoa(respStruct.Code) + ", respStruct.Message: " + respStruct.Message)
	}

	return respStruct.Data.OrderId, nil
}

// 获取dolos配置
func (mpc *MexcPerpetualClient) getDolosConfig() error {
	// https://www.mexc.com/ucgateway/device_api/dolos/config?mhash=b2d52c26ce0879d5d704893fa6e026a6
	// 获取scene 28的数据

	payload := map[string]interface{}{
		"ts":            time.Now().UnixMilli(),
		"type":          0,
		"platform_type": 3,
		"product_type":  0,
		"scene":         0,
		"app_v":         "",
		"sdk_v":         "0.0.13",
		"mtoken":        mpc.mToken,
	}
	reqUrl := "https://www.mexc.com/ucgateway/device_api/dolos/config?mhash=" + mpc.mHash
	headers := map[string]string{
		"Content-Type": "application/json",
		"Origin":       "https://futures.mexc.com/",
		"Referer":      "https://futures.mexc.com/",
		"User-Agent":   utils.USER_AGENT,
	}
	cookies := map[string]string{
		"u_id": mpc.uid,
	}
	respCode, respBody, err := mpc.post(reqUrl, headers, cookies, payload)
	if err != nil {
		fmt.Println("getDolosConfig err:", err)
		return err
	}

	defer respBody.Close()
	if respCode != http.StatusOK {
		fmt.Println("getDolosConfig respCode:", respCode)
		return errors.New("getDolosConfig respCode is " + strconv.Itoa(respCode))
	}
	bodyBytes, err := io.ReadAll(respBody)
	if err != nil {
		fmt.Println("getDolosConfig err:", err)
		return err
	}

	// fmt.Println("getDolosConfig respBody:", string(bodyBytes))
	respStruct := models.MexcCommonResponse[string]{}
	if err := json.Unmarshal(bodyBytes, &respStruct); err != nil {
		fmt.Println("getDolosConfig unmarshall respStruct  err:", err)
		return err
	}
	dolosCipherBytes, err := utils.Base64ToBytes(respStruct.Data)
	if err != nil {
		fmt.Println("getDolosConfig Base64ToBytes err:", err)
		return err
	}
	dolosPlainText, err := utils.Aes256GCMDecrypt(dolosCipherBytes, []byte(utils.MEXC_FUTURES_AES_KEY))
	if err != nil {
		fmt.Println("getDolosConfig Aes256GCMDecrypt err:", err)
		return err
	}
	dolosConfig := models.DolosConfig{}
	if err := json.Unmarshal([]byte(dolosPlainText), &dolosConfig); err != nil {
		fmt.Println("getDolosConfig unmarshall dolosConfig  err:", err)
		return err
	}
	mpc.dolosConfig = &dolosConfig

	return nil
}

// 获取用户信息
func (mpc *MexcPerpetualClient) getCustomerInfo() error {
	// POST https://www.mexc.com/ucenter/api/customer_info
	reqUrl := "https://www.mexc.com/ucenter/api/customer_info"
	cookies := map[string]string{
		"u_id": mpc.uid,
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "https://futures.mexc.com/",
		"Referer":       "https://futures.mexc.com/",
		"ucenter-token": mpc.uid,
		"device-id":     mpc.mToken,
		"User-Agent":    utils.USER_AGENT,
	}
	respCode, respBody, err := mpc.post(reqUrl, headers, cookies, nil)
	if err != nil {
		fmt.Println("getCustomerInfo post err:", err)
		return err
	}

	defer respBody.Close()
	bodyBytes, err := io.ReadAll(respBody)
	if err != nil {
		fmt.Println("getCustomerInfo err:", err)
		return err
	}

	if respCode != http.StatusOK {
		fmt.Printf("getCustomerInfo respCode: %d respBody:xxxxxx %s", respCode, string(bodyBytes))
		return errors.New("getCustomerInfo respCode is " + strconv.Itoa(respCode))
	}

	// fmt.Println("getCustomerInfo respBody:", string(bodyBytes))
	respStruct := models.MexcCommonResponse[models.CustomerInfo]{}
	if err := json.Unmarshal(bodyBytes, &respStruct); err != nil {
		fmt.Println("getCustomerInfo unmarshall respStruct  err:", err)
		return err
	}

	mpc.customerInfo = &respStruct.Data
	return nil
}

func (mpc *MexcPerpetualClient) getTriggerProtect() (string, error) {
	// GET https://futures.mexc.com/api/v1/private/user_pref/get
	reqUrl := "https://futures.mexc.com/api/v1/private/user_pref/get"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Referer":       "https://futures.mexc.com/zh-MY/exchange/AIXBT_USDT?_from=search",
		"authorization": mpc.uid,
		"User-Agent":    utils.USER_AGENT,
	}
	cookies := map[string]string{
		"u_id":                       mpc.uid,
		"mexc_fingerprint_visitorId": mpc.mToken,
	}
	respCode, respBody, err := mpc.get(reqUrl, headers, cookies)
	if err != nil {
		fmt.Println("getTriggerProtect post err:", err)
		return "", err
	}
	defer respBody.Close()
	bodyBytes, err := io.ReadAll(respBody)
	if err != nil {
		fmt.Println("getTriggerProtect err:", err)
		return "", err
	}

	if respCode != http.StatusOK {
		fmt.Printf("getTriggerProtect respCode: %d respBody: %s.\n", respCode, string(bodyBytes))
		return "", errors.New("getTriggerProtect respCode is " + strconv.Itoa(respCode))
	}

	// fmt.Println("getTriggerProtect respBody:", string(bodyBytes))
	respStruct := models.MexcCommonResponse[models.MexcUserPreference]{}
	if err := json.Unmarshal(bodyBytes, &respStruct); err != nil {
		fmt.Println("getTriggerProtect unmarshall respStruct  err:", err)
		return "", err
	}

	return respStruct.Data.TriggerProtect, nil
}

// 返回的ReadCloser要记得关掉
func (mpc *MexcPerpetualClient) get(reqUrl string, headers, cookies map[string]string) (int, io.ReadCloser, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	if cookies == nil {
		cookies = map[string]string{}
	}

	// 1.构造GET请求，拼接Headers和Cookies参数
	getReq, err := fhttp.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return -1, nil, err
	}

	for k, v := range headers {
		getReq.Header.Add(k, v)
	}
	for k, v := range cookies {
		getReq.AddCookie(&fhttp.Cookie{
			Name:  k,
			Value: v,
		})
	}

	// 2.发送请求
	resp, err := mpc.client.Do(getReq)
	if err != nil {
		return -1, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}

// 返回的ReadCloser要记得关掉
func (mpc *MexcPerpetualClient) post(reqUrl string, headers, cookies map[string]string, data map[string]interface{}) (int, io.ReadCloser, error) {
	// post的请求分为两种，一种是表单提交，一种是json提交。
	// 需要通过headers中的Content-Type来区分，如果是application/json，则是json提交，如果是application/x-www-form-urlencoded或者multipart/form-data，则是表单提交。
	// 两种提交方式的区别在于，表单提交的数据格式是key1=value1&key2=value2，而json提交的数据格式是{"key1":"value1","key2":"value2"}。

	if headers == nil {
		headers = map[string]string{}
	}
	if cookies == nil {
		cookies = map[string]string{}
	}
	if data == nil {
		data = map[string]interface{}{}
	}

	// 1. 判断ContentType来使用不同方式构造请求数据
	contentType := headers["Content-Type"]
	bodyStr := ""
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		// 表单提交
		for key, value := range data {
			bodyStr += key + "=" + fmt.Sprintf("%v", value) + "&"
		}

	} else {
		// 非表单提交的都认为是json提交
		bodyBytes, err := json.Marshal(data)
		if err != nil {
			return -1, nil, err
		}
		bodyStr = string(bodyBytes)
	}

	// 2. 构造POST请求体, 并添加请求头和Cookies
	postReq, err := fhttp.NewRequest(http.MethodPost, reqUrl, strings.NewReader(bodyStr))
	if err != nil {
		return -1, nil, err
	}
	for key, value := range headers {
		postReq.Header.Add(key, value)
	}
	for k, v := range cookies {
		postReq.AddCookie(&fhttp.Cookie{
			Name:  k,
			Value: v,
		})
	}

	// 3. 发送请求
	resp, err := mpc.client.Do(postReq)
	if err != nil {
		return -1, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}
