package wordle

import (
	"errors"
	"fmt"
)

func index(runes []rune, x rune) int {
	for i, r := range runes {
		if r == x {
			return i
		}
	}
	return -1
}

func Match(q, a string) (Spots, error) {
	rq := []rune(q)
	ra := []rune(a)
	if len(rq) != len(ra) {
		return nil, fmt.Errorf("length not match between %q and %q", q, a)
	}
	if len(rq) == 0 {
		return nil, errors.New("short word, must be longer than zero")
	}
	rr := make(Spots, len(rq))
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
