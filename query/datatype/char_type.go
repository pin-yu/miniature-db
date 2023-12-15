package datatype

const CharByteLen = 4

type CharType struct {
	arg int
}

func (t *CharType) ByteLen() int {
	return CharByteLen * t.arg
}

func (t *CharType) Argument() int {
	return t.arg
}

func (t *CharType) IsFixedSize() bool {
	return true
}
