package astatine

type Deck struct {
	ID    string `json:"id"`
	*Cards `json:"-"`
}

func NewDeck() *Deck {
	return &Deck{
		ID: newID(),
	}
}

func (d *Deck) With(cards *Cards) *Deck {
	d.Cards = cards
	return d
}

func (d *Deck) String() string {
	return toString(d)
}
