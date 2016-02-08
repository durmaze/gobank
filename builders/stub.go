package builders

import (
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"
)

type StubElement struct {
	Responses  []responses.Response   `json:"responses"`
	Predicates []predicates.Predicate `json:"predicates"`
}

type stubBuilder struct {
	responses  []responses.Response
	predicates []predicates.Predicate
}

type StubBuilder interface {
	Responses(...responses.Response) StubBuilder
	Predicates(...predicates.Predicate) StubBuilder

	Build() StubElement
}

func (builder *stubBuilder) Responses(responses ...responses.Response) StubBuilder {
	builder.responses = responses

	return builder
}

func (builder *stubBuilder) Predicates(predicates ...predicates.Predicate) StubBuilder {
	builder.predicates = predicates

	return builder
}

func (builder *stubBuilder) Build() StubElement {
	return StubElement{
		Responses:  builder.responses,
		Predicates: builder.predicates,
	}
}

func Stub() StubBuilder {
	return &stubBuilder{}
}
