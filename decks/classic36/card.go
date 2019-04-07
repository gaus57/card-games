package classic36

type Card struct {
	cod   string
	name  string
	suit  Suit
	image string
}

func (c *Card) Code() string {
	return c.cod
}

func (c *Card) SuitCode() string {
	return c.suit.cod
}
