package sqltype

type Schema struct {
	fields map[string]Type
}

func NewSchema() *Schema {
	return &Schema{
		fields: make(map[string]Type),
	}
}

func (s *Schema) AddField(fldName string, typ Type) {
	s.fields[fldName] = typ
}
