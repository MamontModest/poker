package players

type PokerGame struct {
	Max_players   int
	Small_blind   int
	Starе_balance int
	Players       []Player
	Game_start    bool
}
type Table map[int]*PokerGame

func NewPokerGame(max int, small int, balance int, status bool) *PokerGame {
	return &PokerGame{
		Max_players:   max,
		Small_blind:   small,
		Starе_balance: balance,
		Game_start:    status,
		Players:       make([]Player, 0),
	}
}

func Connect_game(t *PokerGame, user *Player) {
	t.Players = append(t.Players, *user)

}
