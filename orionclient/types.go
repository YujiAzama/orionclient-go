package orionclient

type Version struct {
	Orion	struct {
		Version	string `json:"version"`
		Uptime	string `json:"uptime"`
		GitHash	string `json:"git_hash"`
		CompileTime	string `json:"compile_time"`
		CompiledBy	string `json:"compiled_by"`
		CompiledIn	string `json:"compiled_in"`
		ReleaseDate	string `json:"release_date"`
		Doc	string `json:"doc"`
	} `json:"orion"`
}

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
