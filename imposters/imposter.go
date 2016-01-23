package imposters

import (
	"fmt"
)

type Imposter struct{
	Protocol string
}

type imposterBuilder struct{
	protocol string
}

type ImposterBuilder interface {
	Protocol(string) ImposterBuilder
	Build() Imposter
}

func (builder *imposterBuilder) Protocol(protocol string) ImposterBuilder {
	builder.protocol = protocol

	return builder
}

func (builder *imposterBuilder) Build() Imposter {
	return Imposter{
		Protocol: builder.protocol,
	}
}

func New() ImposterBuilder {
	return &imposterBuilder{}
}

func Main() {
	fmt.Println("hello")
}