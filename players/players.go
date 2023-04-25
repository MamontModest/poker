package players

type Card struct {
	//масть
	suit string
	//значение
	value string
}

type Cards struct {
	firt_card   *Card
	second_card *Card
}

type Player struct {
	//id player
	uid int
	//карты игрока
	cards *Cards
	//статус в раздаче
	status bool
}

func NewPlayer(uid int) *Player {
	return &Player{
		uid: uid,
		cards: &Cards{
			firt_card: &Card{
				suit:  "no",
				value: "no",
			},
			second_card: &Card{
				suit:  "no",
				value: "no",
			},
		},
		status: true,
	}
}
