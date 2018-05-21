package model

type AuditCollection struct {
	Outstanding []Outstanding `json:"outstanding"`
	MostRecent  []MostRecent  `json:"most-recent"`
}

type Outstanding struct {
	Path        string `json:"path"`
	StartTime   string `json:"start-time"`
	Host        string `json:"host"`
	ElapsedTime string `json:"elapsed-time"`
	Id          string `json:"id"`
}

type MostRecent struct {
	Path        string `json:"path"`
	StartTime   string `json:"start-time"`
	Host        string `json:"host"`
	ElapsedTime string `json:"elapsed-time"`
	Id          string `json:"id"`
}
