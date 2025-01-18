package models

type MexcCommonResponse[T any] struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      T      `json:"data"`
}
