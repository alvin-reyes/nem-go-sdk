package model

type BootNodeRequest struct {
	Metadata struct {
		Application string `json:"application"`
	} `json:"metaData"`
	Endpoint struct {
		Protocol string `json:"protocol"`
		Port     string `json:"port"`
		Host     string `json:"host"`
	} `json:"endpoint"`
	Identity struct {
		PrivateKey string `json:"private-key"`
		Name       string `json:"name"`
	} `json:"identity"`
}
