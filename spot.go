package wordle

import "fmt"

type Spot rune

const (
	NoSpot      Spot = '_'
	WrongSpot        = 'W'
	CorrectSpot      = 'C'
)

type Spots []Spot

func NewSpots(s string) Spots {
	p := Spots(s)
	for i, c := range p {
		switch c {
		case NoSpot, WrongSpot, CorrectSpot:
		default:
			p[i] = NoSpot
		}
	}
	return p
}

func (p Spots) toString() string {
	runes := make([]rune, len(p))
	for i, c := range p {
		runes[i] = rune(c)
	}
	return string(runes)
}

func (p Spots) String() string {
	return p.toString()
}

func (p Spots) GoString() string {
	return p.toString()
}

// Filter creates a Filter from spots and a query word.
func (p Spots) Filter(q string) (Filter, error) {
	runes := []rune(q)
	if len(p) != len(runes) {
		return nil, fmt.Errorf("mismatch length. runes length should be %d", len(p))
	}
	ff := make(filters, 0, len(p))
	cm := make(map[rune][]int)
	for i, sp := range p {
		switch sp {
		case NoSpot:
			ff = append(ff, noSpotFilter(runes[i]))
		case CorrectSpot:
			r := runes[i]
			ff = append(ff, correctSpotFilter(r, i))
			cm[r] = append(cm[r], i)
		}
	}
	for i, sp := range p {
		if sp != WrongSpot {
			continue
		}
		r := runes[i]
		ff = append(ff, wrongSpotFilter(r, i, cm[r]))
	}
	return ff, nil
}
