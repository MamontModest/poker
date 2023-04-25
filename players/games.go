package players

type PokerGame struct {
	max_players  int
	players_now  int
	small_blind  int
	game_start   bool
	star_balance int
	Player       []Player
}
type Tables struct {
	active_tables map[int]PokerGame
}

func NewTablesTable() *Tables {
	return &Tables{
		active_tables: make(map[int]PokerGame),
	}
}
func Ad_table(t *Tables, uid int, p *PokerGame) {
	t.active_tables[uid] = *p
}
func NewPokerGame(max int, now int, small int, balance int, status bool) *PokerGame {
	return &PokerGame{
		max_players:  max,
		players_now:  now,
		small_blind:  small,
		star_balance: balance,
		game_start:   status,
	}
}
