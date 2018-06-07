package astatine

type Note struct {
	ID     string  `json:"id"`
	Fields *Fields `json:"fields"`
}

func NewNote() *Note {
	return &Note{
		ID:     newID(),
		Fields: NewFields(),
	}
}

func (n *Note) With(fields *Fields) *Note {
	n.Fields = fields
	return n
}

func (n *Note) Add(field *Field) *Note {

	n.Fields.Fields = append(n.Fields.Fields, field)
	return n
}

func (n *Note) String() string {
	return toString(n)
}
