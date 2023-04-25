package main

import (
	"fmt"
	"github.com/MamontModest/Poker/players"
)

func main() {
	list_tables := players.NewTablesTable()
	fmt.Println(list_tables)
	var s string
	var uid, room int
	var max, now, small, balance int
	for {
		fmt.Scan(&s)
		if s == "connect" {
			fmt.Scan(&uid, &room)

		} else {
			fmt.Scan(&room, &max, &now, &small, &balance)
			players.Ad_table(list_tables, room, players.NewPokerGame(max, now, small, balance, false))
		}
		fmt.Println(list_tables)
	}

}
