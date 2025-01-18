package services

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/models"
	"github.com/testzhaoxiaofei/mexc-perpetual-api-master/utils"
)

// 生成订单的签名字段
//
// 返回的不包含mToken.mHash,cHash。这三个返回后自己设置进去
func GenerateOrderCreateSig(mToken, memberId string, dolosConfigParams []string) (*models.MexcOrderSig, error) {

	rnd16Bytes, err := utils.GenerateRandomBytes(16)
	if err != nil {
		return nil, err
	}
	len32AesKey := hex.EncodeToString(rnd16Bytes)

	fieldP0, err := calcP0(mToken, memberId, len32AesKey, dolosConfigParams)
	if err != nil {
		return nil, err
	}
	fieldK0, err := calcK0(utils.MEXC_FUTURES_RSA_PUBLIC_KEY, len32AesKey)
	if err != nil {
		return nil, err
	}

	orderSig := models.MexcOrderSig{
		P0: fieldP0,
		K0: fieldK0,
		Ts: time.Now().UnixMilli(),
	}

	return &orderSig, nil
}

// 计算订单签名的p0字段, 返回Base64编码的字符串
//
// @param len32AesKey 32个长度的AesKey字符串
func calcP0(mToken, memberId, len32AesKey string, dolosConfigParams []string) (string, error) {
	browserFp, err := getFingetprint(mToken, memberId)
	if err != nil {
		return "", err
	}
	fingerPrintMap := browserFp.ToMap()
	filteredData := make(map[string]interface{})
	for _, param := range dolosConfigParams {
		if value, ok := fingerPrintMap[param]; ok {
			filteredData[param] = value
		}
	}
	filteredDataBytes, err := json.Marshal(filteredData)
	if err != nil {
		return "", err
	}

	// rnd16Bytes, err := utils.GenerateRandomBytes(16)
	// if err != nil {
	// 	return "", err
	// }
	// rnd32StrAsKey := hex.EncodeToString(rnd16Bytes)
	// 对filteredDataBytes进行AES-256-GCM加密
	encryptedBytes, err := utils.Aes256GCMEncrypt(filteredDataBytes, []byte(len32AesKey))
	if err != nil {
		return "", err
	}

	// 对加密后的字节数组进行Base64编码
	base64EncodedStr := base64.StdEncoding.EncodeToString(encryptedBytes)

	return base64EncodedStr, nil
}

// 计算订单签名的k0字段，返回Base64编码的字符串
func calcK0(rsaPublicKey, plainText string) (string, error) {

	encryptedBytes, err := utils.RSAEncrypt(rsaPublicKey, plainText)
	if err != nil {
		return "", err
	}

	base64EncodedStr := base64.StdEncoding.EncodeToString(encryptedBytes)

	return base64EncodedStr, nil
}

func HeaderSign(cookieUid string, orderMapData map[string]interface{}) (string, string, error) {
	orderDatabytes, err := json.Marshal(orderMapData)
	if err != nil {
		return "", "", err
	}
	orderDataStr := string(orderDatabytes)
	// slog.Info("headerSign", slog.String("orderData", orderDataStr))

	timeInMs := time.Now().UnixMilli()
	timeStr := fmt.Sprintf("%d", timeInMs)
	oMd5 := utils.CalculateMD5(cookieUid + timeStr)
	oMd5Sub := oMd5[7:]
	xMxcSign := utils.CalculateMD5(timeStr + orderDataStr + oMd5Sub)

	return timeStr, xMxcSign, nil
}
