package builders

type Imposter struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Stubs    []stub `json:"stubs"`
}

type ImposterBuilder struct {
	protocol string
	port     int
	stubs    []stub
}

func (builder *ImposterBuilder) Protocol(protocol string) *ImposterBuilder {
	builder.protocol = protocol

	return builder
}

func (builder *ImposterBuilder) Port(port int) *ImposterBuilder {
	builder.port = port

	return builder
}

func (builder *ImposterBuilder) Stubs(stubs ...stub) *ImposterBuilder {
	builder.stubs = stubs

	return builder
}

func (builder *ImposterBuilder) Build() Imposter {
	return Imposter{
		Protocol: builder.protocol,
		Port:     builder.port,
		Stubs:    builder.stubs,
	}
}

func NewImposterBuilder() *ImposterBuilder {
	return &ImposterBuilder{}
}
