package models

import (
	"math"
	"math/rand"
)

type Game struct {
	deck  Deck
	hands map[int][]string
}

func createGame(userIds []int) Game {
	deck := createNewDeck()
	hands := make(map[int][]string)
	empty := []string{}
	for _, user := range userIds {
		hands[user] = empty
	}
	var game Game = Game{
		deck:  deck,
		hands: hands,
	}
	return game
}

func (g *Game) getHand(userId int) []string {
	return g.hands[userId]
}

func (g *Game) swapNewCard(card string, user int, removeIndex int) string {
	hand := g.hands[user]
	toRemove := hand[removeIndex]
	hand[removeIndex] = card
	return toRemove
}

func (g *Game) swapCards(
	user1 int,
	user2 int,
	cardIndex1 int,
	cardIndex2 int,
) {
	hand1 := g.hands[user1]
	hand2 := g.hands[user2]

	hand1[cardIndex1], hand2[cardIndex2] = hand2[cardIndex2], hand1[cardIndex1]

	g.hands[user1] = hand1
	g.hands[user2] = hand2
}

func (g *Game) shuffleHand(user int) {
	hand := g.hands[user]
	for i := 0; i < 3; i++ {
		r := rand.Intn(3)
		r = r % 3
		hand[i], hand[r] = hand[r], hand[i]
	}
}

func (g *Game) getWinner() int {
	winner := 0
	minSum := math.MaxInt
	for key, value := range g.hands {
		sum := 0
		for _, c := range value {
			sum += int(c[0])
		}
		if sum < minSum {
			minSum = sum
			winner = key
		}
	}
	return winner
}
