package astatine

import "github.com/sirupsen/logrus"

type Card struct {
	ID         string       `json:"id"`
	*Notes                  `json:"-"`
	Question   string       `json:"question"`
	Answer     string       `json:"answer"`
	formatFunc func() error `json:"-"`
}

func NewCard() *Card {
	return &Card{
		ID:         newID(),
		formatFunc: func() error { return nil },
	}
}

func (c *Card) WithFunc(formatFunc func() error) *Card {
	c.formatFunc = formatFunc
	return c
}

func (c *Card) Apply() {
	if err := c.formatFunc(); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func (c *Card) With(notes *Notes) *Card {
	c.Notes = notes
	c.Apply()
	return c
}

func (c *Card) Add(note *Note) {
	c.Notes.Notes = append(c.Notes.Notes, note)
	c.Apply()
}

func (c *Card) String() string {
	return toString(c)
}
