package astatine

type Cards struct {
	Cards []*Card `json:"cards"`
}

func NewCards() *Cards {
	return &Cards{}
}

func (c *Cards) Add(card *Card) *Cards {
	if c.Exists(card.ID) {
		return c
	}
	c.Cards = append(c.Cards, card)
	return c
}

func (c *Cards) Get(id string) *Card {
	if c == nil {
		return nil
	}
	for _, card := range c.Cards {
		if card.ID == id {
			return card
		}
	}
	return nil
}

func (c *Cards) String() string {
	return toString(c)
}

func (c *Cards) Exists(id string) bool {
	return c.Get(id) != nil
}
