package datatype

type Type interface {
	ByteLen() int
	Argument() int
	IsFixedSize() bool
}

var _ Type = (*IntType)(nil)
var _ Type = (*CharType)(nil)
