package players

import (
	"errors"
	"fmt"
)

//// тут должна быть логика раунда в игре

type drop struct {
	players     []Player
	max_bet     int
	bets        map[Player]int
	small_blind int
	folt        []Player
}

func create_drop(p []Player, s int) *drop {
	return &drop{
		max_bet:     0,
		players:     p,
		bets:        make(map[Player]int),
		small_blind: s,
		folt:        make([]Player, 0),
	}
}
func Start_round(t *PokerGame) {
	players_club := make([]Player, len(t.Players))
	copy(players_club, t.Players)
	play_drop(create_drop(players_club, t.Small_blind))
}
func play_drop(d *drop) {
	var s string
	var uid int

	d.players[0].Bank -= d.small_blind

	fmt.Println(d.players[0].Bank, "small_blind\n", "cards", d.players[0].Cards.Firt_card, d.players[0].Cards.Second_card)
	d.players[1].Bank -= d.small_blind * 20

	for len(d.folt) != 1 {
		fmt.Scan(&s, &uid)
		switch s {
		case "falt":
			d.falt(uid)
		case "call":
			err := d.call(uid)
			if err != nil {
				fmt.Println("print again bro")
			}

		}
	}
	fmt.Println("fall")
}

func (p *drop) falt(uid int) {
	for _, v := range p.players {
		if v.Uid == uid {
			v.Status = false
			p.folt = append(p.folt, v)
		}
	}
}
func (p *drop) call(uid int) error {
	for _, v := range p.players {
		if v.Uid == uid {
			if v.Bank+p.bets[v] < p.max_bet {
				return errors.New("not enough money")
			}
			p.bets[v] = p.max_bet
			break
		}
	}
	return nil
}
