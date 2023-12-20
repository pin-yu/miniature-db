package sqltype

type Operator interface {
	isSatisfied(lhs Expression, rhs Expression, rec Record) bool
}

type EqualOperator struct {
}

func (o *EqualOperator) isSatisfied(lhs Expression, rhs Expression, rec Record) bool {
	return true
}
