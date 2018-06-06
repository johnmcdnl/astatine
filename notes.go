package astatine

type Notes struct {
	Notes []*Note `json:"notes"`
}

func NewNotes() *Notes {
	return &Notes{}
}

func (n *Notes) Add(note *Note) *Notes {
	if n.Exists(note.ID) {
		return n
	}
	n.Notes = append(n.Notes, note)
	return n
}

func (n *Notes) Get(id string) *Note {
	for _, note := range n.Notes {
		if note.ID == id {
			return note
		}
	}
	return nil
}

func (n *Notes) String() string {
	return toString(n)
}

func (n *Notes) Exists(id string) bool {
	return n.Get(id) != nil
}
