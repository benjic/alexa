package alexa

type Application struct {
	ID string `json:"applicationId"`
}

type RequestContext struct {
	System System `json:"system"`
}

type Permissions struct {
	ConsentToken string `json:"consentToken"`
}

type Session struct {
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	ID          string                 `json:"sessionId"`
	New         bool                   `json:"new"`
	User        User                   `json:"user"`
}

type System struct {
	APIEndpoint string      `json:"apiEndpoint"`
	Application Application `json:"application"`
	User        User        `json:"user"`
}

type User struct {
	AccessToken string      `json:"accessToken"`
	ID          string      `json:"userId"`
	Permissions Permissions `json:"permissions"`
}
