package predicates

type not struct {
	Predicate Predicate `json:"not"`
}

type NotBuilder struct {
	not not
}

func (p not) Type() string {
	return "not"
}

func Not() *NotBuilder {
	return &NotBuilder{not: not{}}
}

func (builder *NotBuilder) Predicate(predicate Predicate) *NotBuilder {
	builder.not.Predicate = predicate

	return builder
}

func (builder *NotBuilder) Build() not {
	return builder.not
}
