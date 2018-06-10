package predicates

import "encoding/json"

type endsWith struct {
	req Request
}

func (p endsWith) Type() string {
	return "endsWith"
}

func (p endsWith) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	endsWithJSON := " { \"endsWith\" : " + requestJSON + "}"

	return []byte(endsWithJSON), nil
}

type EndsWithBuilder struct {
	endsWith endsWith
}

func EndsWith() *EndsWithBuilder {
	return &EndsWithBuilder{endsWith: endsWith{req: Request{}}}
}

func (builder *EndsWithBuilder) Path(path string) *EndsWithBuilder {
	builder.endsWith.req.Path = path
	return builder
}

func (builder *EndsWithBuilder) Method(method string) *EndsWithBuilder {
	builder.endsWith.req.Method = method
	return builder
}

func (builder *EndsWithBuilder) Header(header string, value string) *EndsWithBuilder {
	if builder.endsWith.req.Headers == nil {
		builder.endsWith.req.Headers = map[string]string{}
	}
	builder.endsWith.req.Headers[header] = value
	return builder
}

func (builder *EndsWithBuilder) Query(param string, value string) *EndsWithBuilder {
	if builder.endsWith.req.QueryParams == nil {
		builder.endsWith.req.QueryParams = map[string]string{}
	}

	builder.endsWith.req.QueryParams[param] = value
	return builder
}

func (builder *EndsWithBuilder) Body(body string) *EndsWithBuilder {
	builder.endsWith.req.Body = body
	return builder
}

func (builder *EndsWithBuilder) Build() endsWith {
	return builder.endsWith
}
