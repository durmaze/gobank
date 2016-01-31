package predicates

type or struct {
	Predicates []Predicate `json:"or"`
}

type OrBuilder struct {
	or or
}

func (p or) Type() string {
	return "or"
}

func Or() *OrBuilder {
	return &OrBuilder{or: or{}}
}

func (builder *OrBuilder) Predicates(predicates ...Predicate) *OrBuilder {
	builder.or.Predicates = predicates

	return builder
}

func (builder *OrBuilder) Build() or {
	return builder.or
}
