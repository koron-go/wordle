package wordle_test

import (
	"testing"

	"github.com/koron-go/wordle"
)

func TestMatch(t *testing.T) {
	for _, c := range []struct{ q, a, r string }{
		{"STOTT", "TACIT", "_W__C"},
		// TODO: add cases
	} {
		got, err := wordle.Match(c.q, c.a)
		if err != nil {
			t.Fatalf("wordle.Match failed: %s", err)
		}
		want := wordle.NewSpots(c.r)
		if got.String() != want.String() {
			t.Errorf("unexpected result: q=%s a=%s want=%s got=%s", c.q, c.a, want, got)
		}
	}
}
