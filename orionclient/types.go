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
			Expression *struct {
				Q string `json:"q,omitempty"`
			} `json:"expression,omitempty"`
		} `json:"condition,omitempty"`
	} `json:"subject"`
	Notification struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
		Attrs []string `json:"attrs,omitempty"`
		AttrsFormat string `json:"attrsFormat,omitempty"`
		LastFailure string `json:"lastFailure,omitempty"`
		LastFailureReason string `json:"lastFailureReason,omitempty"`
		LastNotification string `json:"lastNotification,omitempty"`
		LastSuccess string `json:"lastSuccess,omitempty"`
		LastSuccessCode int `json:"lastSuccessCode,omitempty"`
		OnlyChangedAttrs bool `json:"onlyChangedAttrs,omitempty"`
		TimesSent int `json:"timesSent,omitempty"`
	} `json:"notification"`
	Expires    string `json:"expires,omitempty"`
	Throttling int       `json:"throttling,omitempty"`
	Status string `json:"status,omitempty"`
}

type Registration struct {
	Id string `json:"id, omitempty"`
	DataProvided struct {
		Entities []struct {
			ID string `json:"id,omitempty"`
			IdPattern string `"idPattern,omitempty"`
			Type string `json:"type"`
		} `json:"entities"`
		Attrs []string `json:"attrs"`
	} `json:"dataProvided"`
	Provider struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
		LegacyForwarding        bool   `json:"legacyForwarding,omitempty"`
		SupportedForwardingMode string `json:"supportedForwardingMode,omitempty"`
	} `json:"provider"`
	Status string `json:"status,omitempty"`
}
