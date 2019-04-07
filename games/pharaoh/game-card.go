package pharaoh

var cardPoints = map[string]int{
	"ace":   4,
	"king":  3,
	"queen": 30,
	"jack":  2,
	"10":    10,
	"9":     9,
	"8":     8,
	"7":     7,
	"6":     6,
}

type Card interface {
	Code() string
	SuitCode() string
}

type GameCard struct {
	card Card
}

func (gc *GameCard) code() string {
	return gc.card.Code() + "-" + gc.card.SuitCode()
}

func (gc *GameCard) points() int {
	if points, ok := cardPoints[gc.card.Code()]; ok {
		return points
	}
	return 0
}
