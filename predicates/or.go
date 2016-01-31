package predicates

type Or struct {
	Predicates []Predicate `json:"or"`
}

type OrBuilder struct {
	or Or
}

func (p Or) Type() string {
	return "Or"
}

func NewOrBuilder() *OrBuilder {
	return &OrBuilder{or: Or{}}
}

func (builder *OrBuilder) AddPredicate(predicate Predicate) *OrBuilder {
	builder.or.Predicates = append(builder.or.Predicates, predicate)

	return builder
}

func (builder *OrBuilder) Build() Or {
	return builder.or
}
