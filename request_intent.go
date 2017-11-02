package alexa

const (
	intentRequestType = "IntentRequest"
)

type IntentRequest struct {
	Version string            `json:"version"`
	Context RequestContext    `json:"context"`
	Session Session           `json:"session"`
	Request IntentRequestData `json:"request"`
}

type IntentRequestData struct {
	RequestID   string       `json:"requestId"`
	Timestamp   string       `json:"timestamp"`
	DialogState string       `json:"dialogState"`
	Locale      string       `json:"locale"`
	Intent      IntentObject `json:"intent"`
}

type IntentObject struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus"`
	Slots              map[string]SlotObject `json:"slots"`
}

type SlotObject struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
}

type ResolutionsObject struct {
	ResolutionsPerAuthority []ResolutionObject `json:"resolutionsPerAuthority"`
}

type ResolutionObject struct {
	Authority string                  `json:"authority"`
	Status    string                  `json:"status"`
	Values    []ResolutionValueObject `json:"values"`
}

type ResolutionStatus struct {
	Code string `json:"code"`
}

type ResolutionValueObject struct {
	Value ResolutionValue `json:"value"`
}

type ResolutionValue struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
