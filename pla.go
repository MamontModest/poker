package main

import (
	"fmt"
	"github.com/MamontModest/Poker/players"
)

func main() {
	list_tables := make(players.Table)
	var s string
	var uid, room int
	var max, small, balance int
	for {
		fmt.Scan(&s)
		if s == "connect" {
			fmt.Scan(&uid, &room)
			list_tables[room].Players = append(list_tables[room].Players, *players.NewPlayer(uid))
		} else {
			fmt.Scan(&room, &max, &small, &balance)
			list_tables[room] = players.NewPokerGame(max, small, balance, false)
		}
		fmt.Println(list_tables[1].Players)
	}
}
