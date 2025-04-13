package hexagonal

type Struct struct {
	StructName string  `json:"name"`
	Fields     Fields  `json:"fields"`
	Type       string  `json:"type"`
	Children   *Struct `json:"children"`
}

type Field struct {
	FieldName string  `json:"name"`
	FieldType string  `json:"type"`
	Tag       string  `json:"tag"`
	Fields    Fields  `json:"fields"`
	Children  *Struct `json:"children"`
}

type Fields []Field

func (f Fields) AddID() Fields {
	for _, field := range f {
		if field.FieldName == "ID" {
			return f
		}
	}

	newFields := make(Fields, len(f)+1)
	newFields = append(newFields, Field{
		FieldName: "ID",
		FieldType: "string",
		Tag:       "`json:\"id\"`",
	})

	newFields = append(newFields, f...)
	return newFields
}
