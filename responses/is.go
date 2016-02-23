package responses

type IsElement struct {
	Is       *IsResponse `json:"is"`
	Behavior *Behavior   `json:"_behaviors,omitempty"`
}

type IsResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func (i IsElement) Type() string {
	return "is"
}

type IsBuilder struct {
	is         *IsResponse
	waitTime   int
	decorateFn string
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

func (builder *IsBuilder) Wait(waitTime int) *IsBuilder {
	if waitTime > 0 {
		builder.waitTime = waitTime
	}
	return builder
}

func (builder *IsBuilder) Decorate(decorateFn string) *IsBuilder {
	builder.decorateFn = decorateFn
	return builder
}

func (builder *IsBuilder) Build() IsElement {
	is := IsElement{Is: builder.is}
	if builder.waitTime > 0 || len(builder.decorateFn) > 0 {
		is.Behavior = &Behavior{builder.waitTime, builder.decorateFn}
	}
	return is
}

func Is() *IsBuilder {
	return &IsBuilder{is: &IsResponse{}}
}
