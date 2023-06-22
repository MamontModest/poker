package players

import "fmt"

//// тут должна быть логика раунда в игре

type drop struct {
	players []Player
	max_bet int
	//map[uid]bet
	bets map[int]int

	small_blind int
	//[]uid
	Fold               []int
	link_active_player int
	link_diller        int
}

func create_drop(p []Player, s int) *drop {
	return &drop{
		max_bet:            0,
		players:            p,
		bets:               make(map[int]int),
		small_blind:        s,
		Fold:               make([]int, 0),
		link_active_player: 0,
	}
}
func Start_round(d *drop) {
	k := 0
	for _, v := range d.players {
		if v.Status {
			k += 1
		}
	}
	if k < 2 {
		return
	}
	var s string
	var uid, value int
	d.blinds(d.small_blind)

	for _, v := range d.players {
		fmt.Println("bank", v.Bank, "\tcards", v.Cards.Firt_card, v.Cards.Second_card)
	}
	i := d.next_player(d.link_diller + 2)
	for i <= d.link_active_player {
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
		i = d.next_player(i + 1)
	}

	for i, v := range d.players {
		fmt.Println(d.players[i].Bank, " -bank", "\tbet- ", d.bets[v.Uid], d.players[i].Status)
	}
	d.link_diller += 1
}
