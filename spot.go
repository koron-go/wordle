package wordle

type Spot rune

const (
	NoSpot      Spot = '_'
	WrongSpot        = 'W'
	CorrectSpot      = 'C'
)

type Result []Spot

func NewResult(s string) Result {
	r := Result(s)
	for i, c := range r {
		switch c {
		case NoSpot, WrongSpot, CorrectSpot:
		default:
			r[i] = NoSpot
		}
	}
	return r
}

func (r Result) toString() string {
	runes := make([]rune, len(r))
	for i, c := range r {
		runes[i] = rune(c)
	}
	return string(runes)
}

func (r Result) String() string {
	return r.toString()
}

func (r Result) GoString() string {
	return r.toString()
}
