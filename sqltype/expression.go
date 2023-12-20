package sqltype

type Expression interface {
	isField() bool
	isConstant() bool
}

type ConstantExpression struct {
}

func (e *ConstantExpression) isConstant() bool {
	return true
}

func (e *ConstantExpression) isField() bool {
	return false
}

type FieldExpression struct {
}

func (e *FieldExpression) isConstant() bool {
	return false
}

func (e *FieldExpression) isField() bool {
	return true
}
