package models

type CustomerInfo struct {
	MemberID             string `json:"memberId"`
	DigitalID            string `json:"digitalId"`
	CustomerServiceToken string `json:"customerServiceToken"`
}
