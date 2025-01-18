package models

type MexcOrderCreateReqeust struct {
	Symbol        string  `json:"symbol"`
	Side          int     `json:"side"` // 3:开空;1:开多
	OpenType      int     `json:"openType"`
	Type          string  `json:"type"`
	Vol           float64 `json:"vol"`
	Leverage      int     `json:"leverage"`
	MarketCeiling bool    `json:"marketCeiling"`
	PriceProtect  string  `json:"priceProtect"`
	MexcOrderSig          // 签名字段
	Chash         string  `json:"chash"`
	Mtoken        string  `json:"mtoken"`
	Mhash         string  `json:"mhash"`
}

// 抹茶订单签名字段
type MexcOrderSig struct {
	P0 string `json:"p0"`
	K0 string `json:"k0"`
	Ts int64  `json:"ts"`
	// Chash  string `json:"chash"`
	// Mtoken string `json:"mtoken"`
	// Mhash  string `json:"mhash"`
}

type MexcOrderCreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderId string `json:"orderId"`
	} `json:"data"`
}

func (r *MexcOrderCreateReqeust) ToMapData() map[string]interface{} {
	return map[string]interface{}{
		"symbol":        r.Symbol,
		"side":          r.Side,
		"openType":      r.OpenType,
		"type":          r.Type,
		"vol":           r.Vol,
		"leverage":      r.Leverage,
		"marketCeiling": r.MarketCeiling,
		"priceProtect":  r.PriceProtect,
		"p0":            r.P0,
		"k0":            r.K0,
		"chash":         r.Chash,
		"mtoken":        r.Mtoken,
		"ts":            r.Ts,
		"mhash":         r.Mhash,
	}
}
