package wordle

type Spot rune

const (
	NoSpot      Spot = '_'
	WrongSpot        = 'W'
	CorrectSpot      = 'C'
)

type Spots []Spot

func NewSpots(s string) Spots {
	sp := Spots(s)
	for i, c := range sp {
		switch c {
		case NoSpot, WrongSpot, CorrectSpot:
		default:
			sp[i] = NoSpot
		}
	}
	return sp
}

func (sp Spots) toString() string {
	runes := make([]rune, len(sp))
	for i, c := range sp {
		runes[i] = rune(c)
	}
	return string(runes)
}

func (sp Spots) String() string {
	return sp.toString()
}

func (sp Spots) GoString() string {
	return sp.toString()
}
