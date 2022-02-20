package wordle_test

import (
	"testing"

	"github.com/koron-go/wordle"
)

func TestFilter(t *testing.T) {
	for i, c := range []struct {
		spots, query, word string
		want               bool
	}{
		{"W____", "ABCDE", "ABCDE", false},
		{"W____", "ABCDE", "XAXXX", true},
		{"W____", "ABCDE", "XXAXX", true},
		{"W____", "ABCDE", "XXXAX", true},
		{"W____", "ABCDE", "XXXXA", true},
		{"W____", "ABCDE", "AAXXX", false},

		{"C____", "ABCDE", "AXXXX", true},
		{"C____", "ABCDE", "XAXXX", false},
		{"C____", "ABCDE", "XXAXX", false},
		{"C____", "ABCDE", "XXXAX", false},
		{"C____", "ABCDE", "XXXXA", false},
		{"C____", "ABCDE", "AAXXX", true},
	} {
		p := wordle.NewSpots(c.spots)
		f, err := p.Filter(c.query)
		if err != nil {
			t.Fatalf("failed to generate filter at #%d: %#v", i, c)
		}
		got := f.Filter(c.word)
		if got != c.want {
			t.Errorf("unexpected want=%t got=%t at #%d: %#v", c.want, got, i, c)
		}
	}
}
