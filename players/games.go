package players

import (
	"fmt"
	"math/rand"
	"time"
)

type PokerGame struct {
	Max_players   int
	Small_blind   int
	Start_balance int
	Players       []Player
	Game_start    bool
}

func NewPokerGame(max int, small int, balance int, status bool) *PokerGame {
	return &PokerGame{
		Max_players:   max,
		Small_blind:   small,
		Start_balance: balance,
		Game_start:    status,
		Players:       make([]Player, 0, max),
	}
}

type Table map[int]*PokerGame

func Game_start(t *PokerGame) {
	Razdacha(t.Players)
	Start_round(t)
}

func Razdacha(p []Player) {
	cards_now := make([]*Card, 0, 42)
	for _, v := range cards {
		cards_now = append(cards_now, v)
	}
	for _, v := range p {
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(len(cards_now))
		card_1 := cards_now[value]
		cards_now = append(cards_now[:value], cards_now[value+1:]...)

		value = rand.Intn(len(cards_now))
		card_2 := cards_now[value]
		cards_now = append(cards_now[:value], cards_now[value+1:]...)
		v.Cards.Firt_card = card_1
		v.Cards.Second_card = card_2
	}
}
func Start_round(t *PokerGame) {
	for _, v := range t.Players {
		fmt.Println(v.Cards.Firt_card)
		fmt.Println(v.Cards.Second_card)
		fmt.Println(v.Bank)
	}
}
