package builders

type Imposter struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	Stubs    []stub `json:"stubs,omitempty"`
}

type ImposterBuilder struct {
	protocol string
	port     int
	name     string
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

func (builder *ImposterBuilder) Name(name string) *ImposterBuilder {
	builder.name = name

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
		Name:     builder.name,
		Stubs:    builder.stubs,
	}
}

func NewImposterBuilder() *ImposterBuilder {
	return &ImposterBuilder{}
}
