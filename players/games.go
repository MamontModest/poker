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
	Game_start    bool
	Players       []Player
	Game_cards    []Card
	Drop          *drop
}

func NewPokerGame(max int, small int, balance int, status bool) *PokerGame {
	return &PokerGame{
		Max_players:   max,
		Small_blind:   small,
		Start_balance: balance,
		Game_start:    status,
		Players:       make([]Player, 0, max),
		Game_cards:    make([]Card, 0),
	}
}

type Table map[int]*PokerGame

func Game_start(t *PokerGame) {
	i := len(t.Players)
	for i >= 2 {
		one_circle(t)
		i = 0
		for _, v := range t.Players {
			if v.Bank > 0 {
				i += 1
			}
		}
	}
}

func Razdacha(t *PokerGame) {
	cards_now := make([]*Card, 0, 42)
	for _, v := range cards {
		cards_now = append(cards_now, v)
	}
	for _, v := range t.Players {
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
	for i := 0; i < 5; i++ {
		value := rand.Intn(len(cards_now))
		game_card := cards_now[value]
		cards_now = append(cards_now[:value], cards_now[value+1:]...)
		t.Game_cards = append(t.Game_cards, *game_card)
	}
}

func one_circle(t *PokerGame) {
	fmt.Println("circle start")
	t.Drop = create_drop(t.Players, t.Small_blind)
	Razdacha(t)
	Start_round(t.Drop)
	_ = t.Drop
	fmt.Println(t.Game_cards[0], t.Game_cards[1], t.Game_cards[2])
	Start_round(t.Drop)
	fmt.Println(t.Game_cards[3])
	Start_round(t.Drop)
	fmt.Println(t.Game_cards[4])
	Start_round(t.Drop)
	fmt.Println("circle end")
}
