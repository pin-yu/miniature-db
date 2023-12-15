package datatype

type IntConstant struct {
	val int32
}

func (ic *IntConstant) CompareTo(c Constant) (int, error) {
	v, ok := c.(*IntConstant)
	if !ok {
		return 0, ErrIllegalArgument
	}
	if ic.val > v.val {
		return 1, nil
	} else if ic.val == v.val {
		return 0, nil
	} else {
		return -1, nil
	}
}

func (ic *IntConstant) Add(c Constant) (Constant, error) {
	v, ok := c.(*IntConstant)
	if !ok {
		return nil, ErrIllegalArgument
	}

	return &IntConstant{
		val: ic.val + v.val, // don't handle the overflow
	}, nil

}

func (ic *IntConstant) Sub(c Constant) (Constant, error) {
	v, ok := c.(*IntConstant)
	if !ok {
		return nil, ErrIllegalArgument
	}

	return &IntConstant{
		val: ic.val - v.val, // don't handle the overflow
	}, nil
}

func (ic *IntConstant) Mul(c Constant) (Constant, error) {
	v, ok := c.(*IntConstant)
	if !ok {
		return nil, ErrIllegalArgument
	}

	return &IntConstant{
		val: ic.val * v.val, // don't handle the overflow
	}, nil
}

func (ic *IntConstant) Div(c Constant) (Constant, error) {
	v, ok := c.(*IntConstant)
	if !ok {
		return nil, ErrIllegalArgument
	}

	if v.val == 0 {
		return nil, ErrDividedByZero
	}

	return &IntConstant{
		val: ic.val / v.val, // don't handle the overflow
	}, nil
}
