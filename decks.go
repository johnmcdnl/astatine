package astatine

type Decks struct {
	Decks   []*Deck `json:"decks"`
	Primary *Deck   `json:"primary"`
}

func NewDecks() *Decks {
	return &Decks{
		Decks: []*Deck{},
	}
}

func (d *Decks) Add(deck *Deck) *Decks {
	if d.Exists(deck.ID) {
		return d
	}
	d.Decks = append(d.Decks, deck)
	d.SetPrimary(deck)
	return d
}

func (d *Decks) Get(id string) *Deck {
	for _, deck := range d.Decks {
		if deck.ID == id {
			return deck
		}
	}
	return nil
}

func (d *Decks) String() string {
	return toString(d)
}

func (d *Decks) Exists(id string) bool {
	return d.Get(id) != nil
}

func (d *Decks) SetPrimary(deck *Deck) {
	d.Primary = deck
}
