package builders

type Stub struct{
	Responses []Response `json:"responses"`
}

type Response struct{
	Is Is `json:"is"`
}

type Is struct{
	StatusCode int `json:"statusCode"`
}

type stubBuilder struct{
	responses []Response
}

type StubBuilder interface {
	AddResponse(Response) StubBuilder
	Build() Stub
}

func (builder *stubBuilder) AddResponse(response Response) StubBuilder {
	builder.responses = []Response{ response }

	return builder
}

func (builder *stubBuilder) Build() Stub {
	return Stub{
		Responses: builder.responses,
	}
}

func NewStubBuilder() StubBuilder {
	return &stubBuilder{}
}