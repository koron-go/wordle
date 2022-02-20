package wordle

import (
	"errors"
	"fmt"
)

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

func (r Result) String() string {
	return string(r)
}

func (r Result) GoString() string {
	return string(r)
}

func index(runes []rune, x rune) int {
	for i, r := range runes {
		if r == x {
			return i
		}
	}
	return -1
}

func Match(q, a string) (Result, error) {
	rq := []rune(q)
	ra := []rune(a)
	if len(rq) != len(ra) {
		return nil, fmt.Errorf("length not match between %q and %q", q, a)
	}
	if len(rq) == 0 {
		return nil, errors.New("short word, must be longer than zero")
	}
	rr := make(Result, len(rq))
	for i, r := range rq {
		if ra[i] == r {
			rr[i] = CorrectSpot
			ra[i] = 0
		} else {
			rr[i] = NoSpot
		}
	}
	for i := range rr {
		if rr[i] != NoSpot {
			continue
		}
		n := index(ra, rq[i])
		if n < 0 {
			continue
		}
		rr[i] = WrongSpot
		ra[n] = 0
	}
	return rr, nil
}
