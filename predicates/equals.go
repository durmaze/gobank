package predicates

import (
	"encoding/json"
	"log"
)

type Equals struct {
	req Request
}

func (e Equals) MarshalJSON() ([]byte, error) {
	log.Println("marshalling ")
	requestBytes, _ := json.Marshal(e.req)

	requestJson := string(requestBytes)
	equalsJson := " { \"equals\" : " + requestJson + "}"

	return []byte(equalsJson), nil
}

type EqualsBuilder struct {
	equals Equals
}

func NewEqualsBuilder() *EqualsBuilder {
	return &EqualsBuilder{equals: Equals{req: Request{}}}
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

func (builder *EqualsBuilder) Build() Equals {
	return builder.equals
}
