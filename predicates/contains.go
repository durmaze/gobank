package predicates

import (
	"encoding/json"
	"log"
)

type Contains struct {
	req Request
}

func (e Contains) MarshalJSON() ([]byte, error) {
	log.Println("marshalling ")
	requestBytes, _ := json.Marshal(e.req)

	requestJson := string(requestBytes)
	containsJson := " { \"contains\" : " + requestJson + "}"

	return []byte(containsJson), nil
}

type ContainsBuilder struct {
	contains Contains
}

func NewContainsBuilder() *ContainsBuilder {
	return &ContainsBuilder{contains: Contains{req: Request{}}}
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

func (builder *ContainsBuilder) Build() Contains {
	return builder.contains
}
