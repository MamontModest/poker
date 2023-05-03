package players

import (
	"fmt"
)

//// тут должна быть логика раунда в игре

type drop struct {
	players []Player
	max_bet int
	//map[uid]bet
	bets        map[int]int
	small_blind int
	//[]uid
	Fold        []int
	link_player int
}

func create_drop(p []Player, s int) *drop {
	return &drop{
		max_bet:     0,
		players:     p,
		bets:        make(map[int]int),
		small_blind: s,
		Fold:        make([]int, 0),
		link_player: 0,
	}
}
func Start_round(t *PokerGame) {
	players_club := make([]Player, len(t.Players))
	copy(players_club, t.Players)
	play_drop(create_drop(players_club, t.Small_blind))
}
func play_drop(d *drop) {
	var s string
	var uid, value int
	j := d.link_player + 2
	d.blinds(d.small_blind)

	for _, v := range d.players {
		fmt.Println("bank", v.Bank, "\tcards", v.Cards.Firt_card, v.Cards.Second_card)
	}

	for j <= d.link_player {
		i, err := d.next_player(j)
		j = i
		if err != nil {
			fmt.Println("остался один игрок")
			break
		}
		fmt.Println("вводи данные")
		fmt.Scan(&s, &uid)
		switch s {
		case "fold":
			d.fold(uid)

		case "call":
			err := d.call(uid)
			if err != nil {
				fmt.Println("print again bro")
			}

		case "raise":
			fmt.Scan(&value)
			err := d.raise(uid, value)
			if err != nil {
				fmt.Println("no raise bro")
			}

		case "all_in":
			d.all_in(uid)

		case "check":
			err := d.check(uid)
			if err != nil {
				fmt.Println("no chek")
			}
		}

		for _, v := range d.players {
			fmt.Println("bank", v.Bank, "\tcards", v.Cards.Firt_card, v.Cards.Second_card, "\tbet", d.bets[v.Uid])
		}
	}

	for i, v := range d.players {
		fmt.Println(d.players[i].Bank, " -bank", "\tbet- ", d.bets[v.Uid], d.players[i].Status)
	}
	fmt.Println("ROUND END")
}
