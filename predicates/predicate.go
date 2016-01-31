package predicates

type Predicate interface {
	MarshalJSON() ([]byte, error)
}
