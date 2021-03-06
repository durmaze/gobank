package gobank

type ImposterElement struct {
	Protocol       string        `json:"protocol"`
	Port           int           `json:"port,omitempty"`
	Name           string        `json:"name,omitempty"`
	Mode           string        `json:"mode,omitempty"`
	Stubs          []StubElement `json:"stubs,omitempty"`
	RecordRequests bool          `json:"recordRequests,omitempty"`
}

type ImposterBuilder struct {
	protocol       string
	port           int
	name           string
	mode           string
	recordRequests bool
	stubs          []StubElement
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

func (builder *ImposterBuilder) RecordRequests(reply bool) *ImposterBuilder {
	builder.recordRequests = reply

	return builder
}

func (builder *ImposterBuilder) Build() ImposterElement {
	return ImposterElement{
		Protocol:       builder.protocol,
		Port:           builder.port,
		Name:           builder.name,
		Mode:           builder.mode,
		Stubs:          builder.stubs,
		RecordRequests: builder.recordRequests,
	}
}

func NewImposterBuilder() *ImposterBuilder {
	return &ImposterBuilder{}
}
