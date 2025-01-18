package models

type MexcUserPreference struct {
	PositionFundingNotify        string `json:"positionFundingNotify"`
	StopOrderNotify              string `json:"stopOrderNotify"`
	TrackOrderNotify             string `json:"trackOrderNotify"`
	LiquidateWarnNotifyThreshold string `json:"liquidateWarnNotifyThreshold"`
	PositionFundingNotifyRatio   string `json:"positionFundingNotifyRatio"`
	OrderConfirmAgain            string `json:"orderConfirmAgain"`
	UnrealizedPnl                string `json:"unrealizedPnl"`
	PlanOrderNotify              string `json:"planOrderNotify"`
	LiquidateNotify              string `json:"liquidateNotify"`
	PositionViewMode             string `json:"positionViewMode"`
	TriggerProtect               string `json:"triggerProtect"`
	UnfinishedMarketOrderTips    string `json:"unfinishedMarketOrderTips"`
}
