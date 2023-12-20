package sqltype

const IntByteLen = 4

type IntType struct {
}

func (t *IntType) ByteLen() int {
	return IntByteLen
}

func (t *IntType) Argument() int {
	return -1
}

func (t *IntType) IsFixedSize() bool {
	return true
}
