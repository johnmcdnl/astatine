package main

import (
	"github.com/johnmcdnl/astatine"
	"fmt"
)

func main() {
	//
	//var decks = astatine.NewDecks()
	//var cards = astatine.NewCards()
	//var notes = astatine.NewNotes()
	var fields = astatine.NewFields()
	var notes = astatine.NewNotes().Add(astatine.NewNote().With(fields))
	var cards = astatine.NewCards().Add(astatine.NewCard().With(notes))
	var decks = astatine.NewDecks().Add(astatine.NewDeck().With(cards))

	fields.Add(astatine.NewField("front", "hello"))
	fields.Add(astatine.NewField("back", "bonjour"))

	fmt.Println(decks)

	decks2 := astatine.NewDecks().
		Add(astatine.NewDeck().
		With(astatine.NewCards().
		Add(astatine.NewCard().
		With(astatine.NewNotes().
		Add(astatine.NewNote().
		With(astatine.NewFields().
		Add(astatine.NewField("k1", "v1")).
		Add(astatine.NewField("k2", "v2")).
		Add(astatine.NewField("k3", "v3")).
		Add(astatine.NewField("k4", "v4")),
	))))))

	astatine.NewAPI(decks2).Run(3000)

}
