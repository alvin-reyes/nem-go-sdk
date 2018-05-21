package model

type TimeSynchronizationResult struct {
	DateTime          string `json:"dateTime"`
	CurrentTimeOffset string `json:"currentTimeOffset"`
	Change            int    `json:"change"`
}
