package hexagonal

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	FieldName string
	FieldType string
	Tag       string
}
