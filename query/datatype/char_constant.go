package datatype

import "strings"

type CharConstant struct {
	typ Type
	val string
}

func NewCharConstant(typ Type, val string) CharConstant {
	return CharConstant{
		typ: typ,
		val: val,
	}
}

func (cc *CharConstant) CompareTo(c Constant) (int, error) {
	v, ok := c.(*CharConstant)
	if !ok {
		return 0, ErrIllegalArgument
	}

	return strings.Compare(cc.val, v.val), nil
}

func (cc *CharConstant) Add(c Constant) (Constant, error) {
	return nil, ErrUnsupportOperation
}

func (cc *CharConstant) Sub(c Constant) (Constant, error) {
	return nil, ErrUnsupportOperation
}

func (cc *CharConstant) Mul(c Constant) (Constant, error) {
	return nil, ErrUnsupportOperation
}

func (cc *CharConstant) Div(c Constant) (Constant, error) {
	return nil, ErrUnsupportOperation
}
