package builders

type Predicate struct {
	Equals Equals `json:"equals"`
}

type Equals struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	QueryParams map[string]string `json:"query"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
}

type predicateBuilder struct {
	equals *Equals
}

type PredicateBuilder interface {
	EqualsPredicate() PredicateBuilder

	Path(path string) PredicateBuilder
	Method(method string) PredicateBuilder
	Header(string, string) PredicateBuilder
	Query(string, string) PredicateBuilder
	Body(string) PredicateBuilder

	Build() Predicate
}

func (builder *predicateBuilder) EqualsPredicate() PredicateBuilder {
	builder.equals = &Equals{
		"", "", map[string]string{}, map[string]string{}, "",
	}
	return builder
}

func (builder *predicateBuilder) Path(path string) PredicateBuilder {
	builder.equals.Path = path
	return builder
}

func (builder *predicateBuilder) Method(method string) PredicateBuilder {
	builder.equals.Method = method
	return builder
}

func (builder *predicateBuilder) Header(header string, value string) PredicateBuilder {
	builder.equals.Headers[header] = value
	return builder
}

func (builder *predicateBuilder) Query(param string, value string) PredicateBuilder {
	builder.equals.QueryParams[param] = value
	return builder
}

func (builder *predicateBuilder) Body(body string) PredicateBuilder {
	builder.equals.Body = body
	return builder
}

func (builder *predicateBuilder) Build() Predicate {
	return Predicate{
		Equals: *builder.equals,
	}
}

func NewPredicateBuilder() PredicateBuilder {
	return &predicateBuilder{}
}
