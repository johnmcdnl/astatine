package astatine

type Field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewField(key string, value string) *Field {
	return &Field{
		Key:   key,
		Value: value,
	}
}

func (f *Field) String() string {
	return toString(f)
}

func (f *Field) GetValue() string {
	if f == nil {
		return ""
	}

	return f.Value
}
