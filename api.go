package astatine

import (
	"github.com/go-chi/chi"
	"net/http"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/go-chi/chi/middleware"
	"time"
)

type API struct {
	Router chi.Router
	Decks  *Decks
}

func NewAPI(decks *Decks) *API {
	var api API
	api.Decks = decks
	api.NewRouter()
	return &api
}

func (a *API) NewRouter() {
	a.Router = chi.NewRouter()

	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.RealIP)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.Timeout(5 * time.Second))

	a.Router.Mount("/debug", middleware.Profiler())

	a.Router.Get("/decks", a.GetDecks)
	a.Router.Get("/decks/{deckID}", a.GetDeck)

	a.Router.Get("/decks/{deckID}/cards", a.GetCards)
	a.Router.Get("/decks/{deckID}/cards/next", a.GetNextCard)
	a.Router.Get("/decks/{deckID}/cards/{cardID}", a.GetCard)
	a.Router.Get("/cards", a.GetCards)
	a.Router.Get("/cards/next", a.GetNextCard)
	a.Router.Get("/cards/{cardID}", a.GetCard)

	a.Router.Get("/decks/{deckID}/cards/{cardID}/notes", a.GetNotes)
	a.Router.Get("/decks/{deckID}/cards/{cardID}/notes/{noteID}", a.GetNote)
	a.Router.Get("/cards/{cardID}/notes", a.GetNotes)
	a.Router.Get("/cards/{cardID}/notes/{noteID}", a.GetNote)

	a.Router.Get("/decks/{deckID}/cards/{cardID}/notes/{noteID}/fields", a.GetFields)
	a.Router.Get("/cards/{cardID}/notes/{noteID}/fields", a.GetFields)
}

func (a *API) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router)
}

func (a *API) GetDecks(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.Decks.String())
}

func (a *API) GetDeck(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getDeck(r).String())
}

func (a *API) GetCards(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getDeck(r).Cards.String())
}

func (a *API) GetCard(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getCard(r).String())
}

func (a *API) GetNextCard(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getDeck(r).Cards.Next().String())
}

func (a *API) GetNotes(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getCard(r).Notes.String())
}

func (a *API) GetNote(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getNote(r).String())
}

func (a *API) GetFields(w http.ResponseWriter, r *http.Request) {
	a.Write(w, a.getNote(r).Fields.String())
}

func (a *API) Write(w http.ResponseWriter, i string) {
	w.Write([]byte(i))
}

func (a *API) getDeck(r *http.Request) *Deck {
	deckID := chi.URLParam(r, "deckID")
	logrus.Infoln("(a *API) getDeck", deckID)
	if deckID == "" {
		return a.Decks.Primary
	}
	return a.Decks.Get(deckID)
}

func (a *API) getCard(r *http.Request) *Card {
	cardID := chi.URLParam(r, "cardID")
	logrus.Infoln("(a *API) getCard", cardID)
	return a.getDeck(r).Cards.Get(cardID)
}

func (a *API) getNote(r *http.Request) *Note {
	noteID := chi.URLParam(r, "noteID")
	logrus.Infoln("(a *API) getNote", noteID)
	return a.getCard(r).Notes.Get(noteID)
}
