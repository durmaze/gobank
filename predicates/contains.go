package predicates

import "encoding/json"

type contains struct {
	req Request
}

func (p contains) Type() string {
	return "contains"
}

func (p contains) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	containsJSON := " { \"contains\" : " + requestJSON + "}"

	return []byte(containsJSON), nil
}

type ContainsBuilder struct {
	contains contains
}

func Contains() *ContainsBuilder {
	return &ContainsBuilder{contains: contains{req: Request{}}}
}

func (builder *ContainsBuilder) Path(path string) *ContainsBuilder {
	builder.contains.req.Path = path
	return builder
}

func (builder *ContainsBuilder) Method(method string) *ContainsBuilder {
	builder.contains.req.Method = method
	return builder
}

func (builder *ContainsBuilder) Header(header string, value string) *ContainsBuilder {
	if builder.contains.req.Headers == nil {
		builder.contains.req.Headers = map[string]string{}
	}
	builder.contains.req.Headers[header] = value
	return builder
}

func (builder *ContainsBuilder) Query(param string, value string) *ContainsBuilder {
	if builder.contains.req.QueryParams == nil {
		builder.contains.req.QueryParams = map[string]string{}
	}

	builder.contains.req.QueryParams[param] = value
	return builder
}

func (builder *ContainsBuilder) Body(body string) *ContainsBuilder {
	builder.contains.req.Body = body
	return builder
}

func (builder *ContainsBuilder) Build() contains {
	return builder.contains
}
