package sqltype

type Term struct {
	Lhs Expression
	Rhs Expression
	Op  Operator
}
