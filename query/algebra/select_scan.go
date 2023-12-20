package algebra

type SelectScan struct {
}

func (s *SelectScan) BeforeFirst() {

}
func (s *SelectScan) Next() bool {
	return false
}
func (s *SelectScan) Close() {

}
func (s *SelectScan) HasField(fieldName string) bool {
	return false
}
