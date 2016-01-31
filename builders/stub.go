package builders

import . "github.com/durmaze/gobank/predicates"
import . "github.com/durmaze/gobank/responses"

type Stub struct {
	Responses  []Response  `json:"responses"`
	Predicates []Predicate `json:"predicates"`
}

type stubBuilder struct {
	responses  []Response
	predicates []Predicate
}

type StubBuilder interface {
	Responses(...Response) StubBuilder
	Predicates(...Predicate) StubBuilder

	Build() Stub
}

func (builder *stubBuilder) Responses(responses ...Response) StubBuilder {
	builder.responses = responses

	return builder
}

func (builder *stubBuilder) Predicates(predicates ...Predicate) StubBuilder {
	builder.predicates = predicates

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
