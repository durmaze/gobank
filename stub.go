package gobank

import (
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"
)

type StubElement struct {
	Responses  []responses.Response   `json:"responses"`
	Predicates []predicates.Predicate `json:"predicates"`
}

type StubBuilder struct {
	responses  []responses.Response
	predicates []predicates.Predicate
}

func (builder *StubBuilder) Responses(responses ...responses.Response) *StubBuilder {
	builder.responses = responses

	return builder
}

func (builder *StubBuilder) Predicates(predicates ...predicates.Predicate) *StubBuilder {
	builder.predicates = predicates

	return builder
}

func (builder *StubBuilder) Build() StubElement {
	return StubElement{
		Responses:  builder.responses,
		Predicates: builder.predicates,
	}
}

func Stub() *StubBuilder {
	return &StubBuilder{}
}
