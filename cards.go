package astatine

import (
	"math/rand"
)

type Cards struct {
	Cards []*Card `json:"cards"`
}

func NewCards() *Cards {
	return &Cards{
		Cards: []*Card{},
	}
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

func (c *Cards) Next() *Card {
	if len(c.Cards) == 0 {
		return nil
	}
	c.Shuffle()
	return c.Cards[0]
}

func (c *Cards) Shuffle() {
	rand.Shuffle(len(c.Cards), func(i, j int) {
		c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
	})
}

func (c *Cards) String() string {
	return toString(c)
}

func (c *Cards) Exists(id string) bool {
	return c.Get(id) != nil
}
