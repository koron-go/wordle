package word_test

import (
	"testing"

	"github.com/koron-go/wordle/word"
)

func TestCount(t *testing.T) {
	if len(word.Answer) == 0 {
		t.Fatal("word.Answer is empty")
	}
	if len(word.Other) == 0 {
		t.Fatal("word.Other is empty")
	}
}

func TestAnswerNotDuplicate(t *testing.T) {
	seen := make(map[string]struct{}, len(word.Answer))
	for _, w := range word.Answer {
		if _, has := seen[w]; has {
			t.Errorf("word %q is duplicated", w)
		}
		seen[w] = struct{}{}
	}
}

func TestOtherNotDuplicate(t *testing.T) {
	seen := make(map[string]struct{}, len(word.Other))
	for _, w := range word.Other {
		if _, has := seen[w]; has {
			t.Errorf("word %q is duplicated", w)
		}
		seen[w] = struct{}{}
	}
}

func TestNotIntersect(t *testing.T) {
	seen := make(map[string]struct{}, len(word.Answer))
	for _, w := range word.Answer {
		seen[w] = struct{}{}
	}
	for _, w := range word.Other {
		if _, has := seen[w]; has {
			t.Errorf("word %q intersect both answer and other sets", w)
		}
	}
}

func TestAllSize(t *testing.T) {
	got := len(word.All)
	want := len(word.Other) + len(word.Answer)
	if got != want {
		t.Errorf("unexpected len(word.All) want=%d got=%d", want, got)
	}
}

func TestAllNotDuplicate(t *testing.T) {
	seen := make(map[string]struct{}, len(word.All))
	for _, w := range word.All {
		if _, has := seen[w]; has {
			t.Errorf("word %q is duplicated", w)
		}
		seen[w] = struct{}{}
	}
}
