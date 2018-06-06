package astatine

type Note struct {
	ID     string  `json:"id"`
	Fields *Fields `json:"fields"`
}

func NewNote() *Note {
	return &Note{
		ID: newID(),
	}
}

func (n *Note) With(fields *Fields) *Note {
	n.Fields = fields
	return n
}

func (n *Note) String() string {
	return toString(n)
}
