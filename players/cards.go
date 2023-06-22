package players

type Card struct {
	//масть
	Suit string
	//значение
	Value string
}

func NewCard(suit string, value string) *Card {
	return &Card{
		Suit:  suit,
		Value: value,
	}
}

var cards = []*Card{
	NewCard("clubs", "ace"), NewCard("diamonds", "ace"), NewCard("hearts", "ace"), NewCard("spade", "ace"),
	NewCard("clubs", "king"), NewCard("diamonds", "king"), NewCard("hearts", "king"), NewCard("spade", "king"),
	NewCard("clubs", "queen"), NewCard("diamonds", "queen"), NewCard("hearts", "queen"), NewCard("spade", "queen"),
	NewCard("clubs", "jack"), NewCard("diamonds", "jack"), NewCard("hearts", "jack"), NewCard("spade", "jack"),
	NewCard("clubs", "ten"), NewCard("diamonds", "ten"), NewCard("hearts", "ten"), NewCard("spade", "ten"),
	NewCard("clubs", "nine"), NewCard("diamonds", "nine"), NewCard("hearts", "nine"), NewCard("spade", "nine"),
	NewCard("clubs", "eight"), NewCard("diamonds", "eight"), NewCard("hearts", "eight"), NewCard("spade", "eight"),
	NewCard("clubs", "seven"), NewCard("diamonds", "seven"), NewCard("hearts", "seven"), NewCard("spade", "seven"),
	NewCard("clubs", "six"), NewCard("diamonds", "six"), NewCard("hearts", "six"), NewCard("spade", "six"),
	NewCard("clubs", "five"), NewCard("diamonds", "five"), NewCard("hearts", "five"), NewCard("spade", "five"),
	NewCard("clubs", "four"), NewCard("diamonds", "four"), NewCard("hearts", "four"), NewCard("spade", "four"),
	NewCard("clubs", "three"), NewCard("diamonds", "three"), NewCard("hearts", "three"), NewCard("spade", "three"),
	NewCard("clubs", "two"), NewCard("diamonds", "two"), NewCard("hearts", "two"), NewCard("spade", "two"),
}

type Cards struct {
	Firt_card   *Card
	Second_card *Card
}
