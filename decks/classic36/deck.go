package classic36

func Suits() [4]Suit {
	return [4]Suit{
		Suit{cod: "diamonds", name: "Бубны", image: ""},
		Suit{cod: "hearts", name: "Червы", image: ""},
		Suit{cod: "spades", name: "Пики", image: ""},
		Suit{cod: "clubs", name: "Трефы", image: ""},
	}
}

func Deck() [36]*Card {
	var dec [36]*Card
	for i, suit := range Suits() {
		dec[0+i*9] = &Card{cod: "6", name: "6", suit: suit, image: ""}
		dec[1+i*9] = &Card{cod: "7", name: "7", suit: suit, image: ""}
		dec[2+i*9] = &Card{cod: "8", name: "8", suit: suit, image: ""}
		dec[3+i*9] = &Card{cod: "9", name: "9", suit: suit, image: ""}
		dec[4+i*9] = &Card{cod: "10", name: "10", suit: suit, image: ""}
		dec[5+i*9] = &Card{cod: "jack", name: "Валет", suit: suit, image: ""}
		dec[6+i*9] = &Card{cod: "queen", name: "Дама", suit: suit, image: ""}
		dec[7+i*9] = &Card{cod: "king", name: "Король", suit: suit, image: ""}
		dec[8+i*9] = &Card{cod: "ace", name: "Туз", suit: suit, image: ""}
	}

	return dec
}
