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
		switch s {
		case "connect":
			fmt.Scan(&room, &uid)
			players.Connect_game(list_tables[room], players.NewPlayer(uid))

			if len(list_tables[room].Players) == list_tables[room].Max_players {
				players.Game_start(list_tables[room])
			}
		case "create":
			fmt.Scan(&room, &max, &small, &balance)
			list_tables[room] = players.NewPokerGame(max, small, balance, false)

		case "leave":
			fmt.Scan(&room, &uid)
			for i, v := range list_tables[room].Players {
				if v.Uid == uid {
					list_tables[room].Players = append(list_tables[room].Players[:i], list_tables[room].Players[i+1:]...)
					break
				}
			}
		}
		/*   Вывод таблицы
		for i, v := range list_tables {
			fmt.Println("room id: ", i, v.Players)
		}
		*/
	}
}
