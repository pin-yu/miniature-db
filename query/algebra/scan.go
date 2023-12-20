package algebra

type Scan interface {
	BeforeFirst()
	Next() bool
	Close()
	HasField(fieldName string) bool
}
