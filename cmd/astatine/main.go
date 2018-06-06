package main

import (
	"github.com/johnmcdnl/astatine"
	"fmt"
)

func main() {

	var decks = astatine.NewDecks()
	var deck = astatine.NewDeck()

	var cards = astatine.NewCards()
	var card = astatine.NewCard()

	var notes = astatine.NewNotes()
	var note = astatine.NewNote()

	var fields = astatine.NewFields()

	decks.Add(deck.With(cards))
	cards.Add(card.With(notes))
	notes.Add(note.With(fields))

	fields.Add(astatine.NewField("front", "hello"))
	fields.Add(astatine.NewField("back", "bonjour"))

	fmt.Println(decks)

	astatine.NewAPI(decks).Run(3000)

}
