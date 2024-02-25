package models

import "math/rand"

type Deck struct {
	Cards []string
}

func createNewDeck() Deck {
	suits := []string{"H", "C", "D", "S"}
	cards := []string{}
	for _, suit := range suits {
		for i := 1; i <= 13; i++ {
			cards = append(cards, suit+string(i))
		}
	}
	var d Deck = Deck{Cards: cards}
	return d
}

func (d *Deck) shuffleDeck() {
	for i := 0; i < 52; i++ {
		r := rand.Intn(52)
		d.Cards[i], d.Cards[r] = d.Cards[r], d.Cards[i]
	}
}

func (d *Deck) pickTopCard() string {
	card := d.Cards[0]
	d.removeCardFromDeck(card)
	return card
}

func (d *Deck) removeCardFromDeck(card string) {
	i := 0
	for i = 0; i < len(d.Cards); i++ {
		if d.Cards[i] == card {
			break
		}
	}
	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
}
