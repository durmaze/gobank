package builders

type Imposter struct{
	Protocol string `json:"protocol"`
	Port int `json:"port"`
}

type imposterBuilder struct{
	protocol string
	port int
}

type ImposterBuilder interface {
	Protocol(string) ImposterBuilder
	Port(int) ImposterBuilder
	Build() Imposter
}

func (builder *imposterBuilder) Protocol(protocol string) ImposterBuilder {
	builder.protocol = protocol

	return builder
}

func (builder *imposterBuilder) Port(port int) ImposterBuilder {
	builder.port = port

	return builder
}

func (builder *imposterBuilder) Build() Imposter {
	return Imposter{
		Protocol: builder.protocol,
		Port: builder.port,
	}
}

func NewImposterBuilder() ImposterBuilder {
	return &imposterBuilder{}
}
