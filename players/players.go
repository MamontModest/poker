package players

type Player struct {
	//id player
	Uid int
	//карты игрока
	Cards *Cards
	//статус в раздаче
	Status bool
	//банк
	Bank int
}

func NewPlayer(uid int) *Player {
	return &Player{
		Uid: uid,
		Cards: &Cards{
			Firt_card:   &Card{},
			Second_card: &Card{},
		},
		Status: true,
		Bank:   0,
	}
}
func Connect_game(t *PokerGame, user *Player) {
	user.Bank = t.Start_balance
	t.Players = append(t.Players, *user)
}
