package model

type NemRequestResult struct {
	Type    int    `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
