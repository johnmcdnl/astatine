package main

import (
	"github.com/johnmcdnl/astatine"
	"github.com/johnmcdnl/astatine/data"
)

func main() {
	astatine.NewAPI(
		astatine.NewDecks().
			Add(data.Load(data.Countries)),
	).Run(3000)
}
