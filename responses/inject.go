package responses

type inject struct {
	InjectedFunction string `json:"inject"`
}

func (i inject) Type() string {
	return "inject"
}

type InjectBuilder struct {
	functionToInject string
}

func (builder *InjectBuilder) Fn(functionCode string) *InjectBuilder {
	builder.functionToInject = functionCode
	return builder
}

func (builder *InjectBuilder) Build() Response {
	return inject{InjectedFunction: builder.functionToInject}
}

func Inject() *InjectBuilder {
	return &InjectBuilder{}
}
