package builders

type Stub struct {
	Responses  []Response  `json:"responses"`
	Predicates []Predicate `json:"predicates"`
}

type stubBuilder struct {
	responses  []Response
	predicates []Predicate
}

type StubBuilder interface {
	AddResponse(Response) StubBuilder
	AddPredicate(Predicate) StubBuilder

	Build() Stub
}

func (builder *stubBuilder) AddResponse(response Response) StubBuilder {
	builder.responses = append(builder.responses, response)

	return builder
}

func (builder *stubBuilder) AddPredicate(predicate Predicate) StubBuilder {
	builder.predicates = append(builder.predicates, predicate)

	return builder
}

func (builder *stubBuilder) Build() Stub {
	return Stub{
		Responses:  builder.responses,
		Predicates: builder.predicates,
	}
}

func NewStubBuilder() StubBuilder {
	return &stubBuilder{}
}
