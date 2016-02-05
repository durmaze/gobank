package predicates

type and struct {
	Predicates []Predicate `json:"and"`
}

type AndBuilder struct {
	and and
}

func (p and) Type() string {
	return "and"
}

func And() *AndBuilder {
	return &AndBuilder{and: and{}}
}

func (builder *AndBuilder) Predicates(predicates ...Predicate) *AndBuilder {
	builder.and.Predicates = predicates

	return builder
}

func (builder *AndBuilder) Build() and {
	return builder.and
}
