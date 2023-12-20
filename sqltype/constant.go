package sqltype

import "errors"

var ErrIllegalArgument = errors.New("illegal argument")
var ErrUnsupportOperation = errors.New("unsupport operation")
var ErrDividedByZero = errors.New("divided by zero")

type Constant interface {
	CompareTo(Constant) (int, error)
	Add(Constant) (Constant, error)
	Sub(Constant) (Constant, error)
	Mul(Constant) (Constant, error)
	Div(Constant) (Constant, error)
}

var _ Constant = (*IntConstant)(nil)
