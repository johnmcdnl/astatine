package astatine

type Deck struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	*Cards      `json:"-"`
}

func NewDeck() *Deck {
	return &Deck{
		ID:    newID(),
		Cards: NewCards(),
	}
}

func (d *Deck) With(cards *Cards) *Deck {
	d.Cards = cards
	return d
}

func (d *Deck) String() string {
	return toString(d)
}
