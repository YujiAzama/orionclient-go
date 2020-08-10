package orionclient

type Subscription struct {
	Id string `json:"id,omitempty"`
	Description string `json:"description"`
	Subject     struct {
		Entities []struct {
			ID   string `json:"id,omitempty"`
			IdPattern string `json:"idPattern,omitempty"`
			Type string `json:"type"`
		} `json:"entities"`
		Condition *struct {
			Attrs []string `json:"attrs,omitempty"`
		} `json:"condition,omitempty"`
	} `json:"subject"`
	Notification struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
		Attrs []string `json:"attrs,omitempty"`
	} `json:"notification"`
	Expires    string `json:"expires,omitempty"`
	Throttling int       `json:"throttling,omitempty"`
}
