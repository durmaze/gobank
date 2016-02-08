package predicates

type inject struct {
	Function string `json:"inject"`
}

func (p inject) Type() string {
	return "inject"
}

type InjectBuilder struct {
	inject inject
}

func Inject() *InjectBuilder {
	return &InjectBuilder{inject: inject{}}
}

func (builder *InjectBuilder) Function(function string) *InjectBuilder {
	builder.inject.Function = function
	return builder
}

func (builder *InjectBuilder) Build() inject {
	return builder.inject
}
