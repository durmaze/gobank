package predicates

type Request struct {
	Method      string            `json:"method,omitempty"`
	Path        string            `json:"path,omitempty"`
	QueryParams map[string]string `json:"query,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Body        string            `json:"body,omitempty"`
}
