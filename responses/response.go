package responses

type Response struct {
	Is Is `json:"is"`
}

type Is struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type responseBuilder struct {
	is *Is
}

type ResponseBuilder interface {
	IsResponse() ResponseBuilder
	StatusCode(int) ResponseBuilder
	Header(string, string) ResponseBuilder
	Body(string) ResponseBuilder
	Build() Response
}

func (builder *responseBuilder) IsResponse() ResponseBuilder {
	builder.is = &Is{
		404, map[string]string{}, "",
	}
	return builder
}

func (builder *responseBuilder) StatusCode(statusCode int) ResponseBuilder {
	builder.is.StatusCode = statusCode
	return builder
}

func (builder *responseBuilder) Header(header string, value string) ResponseBuilder {
	builder.is.Headers[header] = value
	return builder
}

func (builder *responseBuilder) Body(body string) ResponseBuilder {
	builder.is.Body = body
	return builder
}

func (builder *responseBuilder) Build() Response {
	return Response{
		Is: *builder.is,
	}
}

func NewResponseBuilder() ResponseBuilder {
	return &responseBuilder{}
}
