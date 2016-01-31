package responses

type is struct {
	Is isResponse `json:"is"`
}

type isResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func (i is) Type() string {
	return "is"
}

type IsBuilder struct {
	is isResponse
}

func (builder *IsBuilder) StatusCode(statusCode int) *IsBuilder {
	builder.is.StatusCode = statusCode
	return builder
}

func (builder *IsBuilder) Header(header string, value string) *IsBuilder {
	if builder.is.Headers == nil {
		builder.is.Headers = map[string]string{}
	}

	builder.is.Headers[header] = value
	return builder
}

func (builder *IsBuilder) Body(body string) *IsBuilder {
	builder.is.Body = body
	return builder
}

func (builder *IsBuilder) Build() is {
	return is{Is: builder.is}
}

func Is() *IsBuilder {
	return &IsBuilder{is: isResponse{}}
}
