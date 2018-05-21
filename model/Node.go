package model

type Node struct {
	Metadata struct {
		Features    int    `json:"features"`
		NetworkId   uint   `json:"networkId"`
		Application string `json:"application"`
		Version     string `json:"version"`
		Platform    string `json:"platform"`
	}

	Endpoint struct {
		Protocol string `json:"protocol"`
		Port     string `json:"port"`
		Host     string `json:"host"`
	} `json:"endpoint"`
	Identity struct {
		Name      string `json:"name"`
		PublicKey string `json:"public-key"`
	} `json:"identity"`
}
