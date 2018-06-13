package predicates

import "encoding/json"

type equals struct {
	req Request
}

func (p equals) Type() string {
	return "equals"
}

func (p equals) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	equalsJSON := " { \"equals\" : " + requestJSON + "}"

	return []byte(equalsJSON), nil
}

type EqualsBuilder struct {
	equals equals
}

func Equals() *EqualsBuilder {
	return &EqualsBuilder{equals: equals{req: Request{}}}
}

func (builder *EqualsBuilder) Path(path string) *EqualsBuilder {
	builder.equals.req.Path = path
	return builder
}

func (builder *EqualsBuilder) Method(method string) *EqualsBuilder {
	builder.equals.req.Method = method
	return builder
}

func (builder *EqualsBuilder) Header(header string, value string) *EqualsBuilder {
	if builder.equals.req.Headers == nil {
		builder.equals.req.Headers = map[string]string{}
	}
	builder.equals.req.Headers[header] = value
	return builder
}

func (builder *EqualsBuilder) Query(param string, value string) *EqualsBuilder {
	if builder.equals.req.QueryParams == nil {
		builder.equals.req.QueryParams = map[string]string{}
	}

	builder.equals.req.QueryParams[param] = value
	return builder
}

func (builder *EqualsBuilder) Body(body string) *EqualsBuilder {
	builder.equals.req.Body = body
	return builder
}

func (builder *EqualsBuilder) Build() equals {
	return builder.equals
}
