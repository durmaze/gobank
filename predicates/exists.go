package predicates

import "encoding/json"

type existsRequest struct {
	Method      *bool           `json:"method,omitempty"`
	Path        *bool           `json:"path,omitempty"`
	QueryParams map[string]bool `json:"query,omitempty"`
	Headers     map[string]bool `json:"headers,omitempty"`
	Body        *bool           `json:"body,omitempty"`
}

type exists struct {
	req existsRequest
}

func (p exists) Type() string {
	return "exists"
}

func (p exists) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	existsJSON := " { \"exists\" : " + requestJSON + "}"

	return []byte(existsJSON), nil
}

type ExistsBuilder struct {
	exists exists
}

func Exists() *ExistsBuilder {
	return &ExistsBuilder{exists: exists{req: existsRequest{}}}
}

func (builder *ExistsBuilder) Path(path bool) *ExistsBuilder {
	builder.exists.req.Path = &path
	return builder
}

func (builder *ExistsBuilder) Method(method bool) *ExistsBuilder {
	builder.exists.req.Method = &method
	return builder
}

func (builder *ExistsBuilder) Header(header string, value bool) *ExistsBuilder {
	if builder.exists.req.Headers == nil {
		builder.exists.req.Headers = map[string]bool{}
	}
	builder.exists.req.Headers[header] = value
	return builder
}

func (builder *ExistsBuilder) Query(param string, value bool) *ExistsBuilder {
	if builder.exists.req.QueryParams == nil {
		builder.exists.req.QueryParams = map[string]bool{}
	}

	builder.exists.req.QueryParams[param] = value
	return builder
}

func (builder *ExistsBuilder) Body(body bool) *ExistsBuilder {
	builder.exists.req.Body = &body
	return builder
}

func (builder *ExistsBuilder) Build() exists {
	return builder.exists
}
