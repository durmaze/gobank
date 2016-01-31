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

func NewOrBuilder() *OrBuilder {
	return &OrBuilder{or: or{}}
}

func (builder *OrBuilder) AddPredicate(predicate Predicate) *OrBuilder {
	builder.or.Predicates = append(builder.or.Predicates, predicate)

	return builder
}

func (builder *OrBuilder) Build() or {
	return builder.or
}
