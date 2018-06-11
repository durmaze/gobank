package predicates

import "encoding/json"

type startsWith struct {
	req Request
}

func (p startsWith) Type() string {
	return "startsWith"
}

func (p startsWith) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	startsWithJSON := " { \"startsWith\" : " + requestJSON + "}"

	return []byte(startsWithJSON), nil
}

type StartsWithBuilder struct {
	startsWith startsWith
}

func StartsWith() *StartsWithBuilder {
	return &StartsWithBuilder{startsWith: startsWith{req: Request{}}}
}

func (builder *StartsWithBuilder) Path(path string) *StartsWithBuilder {
	builder.startsWith.req.Path = path
	return builder
}

func (builder *StartsWithBuilder) Method(method string) *StartsWithBuilder {
	builder.startsWith.req.Method = method
	return builder
}

func (builder *StartsWithBuilder) Header(header string, value string) *StartsWithBuilder {
	if builder.startsWith.req.Headers == nil {
		builder.startsWith.req.Headers = map[string]string{}
	}
	builder.startsWith.req.Headers[header] = value
	return builder
}

func (builder *StartsWithBuilder) Query(param string, value string) *StartsWithBuilder {
	if builder.startsWith.req.QueryParams == nil {
		builder.startsWith.req.QueryParams = map[string]string{}
	}

	builder.startsWith.req.QueryParams[param] = value
	return builder
}

func (builder *StartsWithBuilder) Body(body string) *StartsWithBuilder {
	builder.startsWith.req.Body = body
	return builder
}

func (builder *StartsWithBuilder) Build() startsWith {
	return builder.startsWith
}
