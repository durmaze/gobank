package predicates

import "encoding/json"

type deepEquals struct {
	req Request
}

func (p deepEquals) Type() string {
	return "deepEquals"
}

func (p deepEquals) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	deepEqualsJSON := " { \"deepEquals\" : " + requestJSON + "}"

	return []byte(deepEqualsJSON), nil
}

type DeepEqualsBuilder struct {
	deepEquals deepEquals
}

func DeepEquals() *DeepEqualsBuilder {
	return &DeepEqualsBuilder{deepEquals: deepEquals{req: Request{}}}
}

func (builder *DeepEqualsBuilder) Path(path string) *DeepEqualsBuilder {
	builder.deepEquals.req.Path = path
	return builder
}

func (builder *DeepEqualsBuilder) Method(method string) *DeepEqualsBuilder {
	builder.deepEquals.req.Method = method
	return builder
}

func (builder *DeepEqualsBuilder) Header(header string, value string) *DeepEqualsBuilder {
	if builder.deepEquals.req.Headers == nil {
		builder.deepEquals.req.Headers = map[string]string{}
	}
	builder.deepEquals.req.Headers[header] = value
	return builder
}

func (builder *DeepEqualsBuilder) Query(param string, value string) *DeepEqualsBuilder {
	if builder.deepEquals.req.QueryParams == nil {
		builder.deepEquals.req.QueryParams = map[string]string{}
	}

	builder.deepEquals.req.QueryParams[param] = value
	return builder
}

func (builder *DeepEqualsBuilder) Body(body string) *DeepEqualsBuilder {
	builder.deepEquals.req.Body = body
	return builder
}

func (builder *DeepEqualsBuilder) Build() deepEquals {
	return builder.deepEquals
}
