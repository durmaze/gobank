package builders

import . "github.com/durmaze/gobank/predicates"
import . "github.com/durmaze/gobank/responses"

type stub struct {
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

	Build() stub
}

func (builder *stubBuilder) Responses(responses ...Response) StubBuilder {
	builder.responses = responses

	return builder
}

func (builder *stubBuilder) Predicates(predicates ...Predicate) StubBuilder {
	builder.predicates = predicates

	return builder
}

func (builder *stubBuilder) Build() stub {
	return stub{
		Responses:  builder.responses,
		Predicates: builder.predicates,
	}
}

func Stub() StubBuilder {
	return &stubBuilder{}
}
