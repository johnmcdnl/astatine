package astatine

type Fields struct {
	Fields []*Field `json:"fields"`
}

func NewFields() *Fields {
	return &Fields{}
}

func (f *Fields) Add(field *Field) {
	if f.Exists(field.Key) {
		return
	}
	f.Fields = append(f.Fields, field)
}

func (f *Fields) Get(key string) *Field {
	for _, field := range f.Fields {
		if field.Key == key {
			return field
		}
	}
	return nil
}

func (f *Fields) String() string {
	return toString(f)
}

func (f *Fields) Exists(key string) bool {
	return f.Get(key) != nil
}
