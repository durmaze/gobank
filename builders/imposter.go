package builders

type Imposter struct {
	Protocol string        `json:"protocol"`
	Port     int           `json:"port,omitempty"`
	Name     string        `json:"name,omitempty"`
	Mode     string        `json:"mode,omitempty"`
	Stubs    []StubElement `json:"stubs,omitempty"`
}

type ImposterBuilder struct {
	protocol string
	port     int
	name     string
	mode     string
	stubs    []StubElement
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

func (builder *ImposterBuilder) Mode(mode string) *ImposterBuilder {
	builder.mode = mode

	return builder
}

func (builder *ImposterBuilder) Stubs(stubs ...StubElement) *ImposterBuilder {
	builder.stubs = stubs

	return builder
}

func (builder *ImposterBuilder) Build() Imposter {
	return Imposter{
		Protocol: builder.protocol,
		Port:     builder.port,
		Name:     builder.name,
		Mode:     builder.mode,
		Stubs:    builder.stubs,
	}
}

func NewImposterBuilder() *ImposterBuilder {
	return &ImposterBuilder{}
}
