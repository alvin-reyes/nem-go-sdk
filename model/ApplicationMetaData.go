package model

type ApplicationMetaData struct {
	CurrentTime string `json:"currentTime"`
	Application string `json:"application"`
	StartTime   string `json:"startTime"`
	Version     int    `json:"version"`
	Signer      string `json:"signer"`
}
