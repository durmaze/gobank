package responses

type Behavior struct {
	Wait       int    `json:"wait,omitempty"`
	DecorateFn string `json:"decorate,omitempty"`
}
