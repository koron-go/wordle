package word

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"
)

//go:embed answer.txt
var answer []byte

//go:embed other.txt
var other []byte

var Answer []string

var Other []string

func load(b []byte) ([]string, error) {
	words := make([]string, 0, len(b)/6)
	r := bufio.NewReader(bytes.NewReader(b))
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return words, nil
			}
			return nil, err
		}
		w := strings.TrimSpace(l)
		if len(w) != 5 {
			return nil, fmt.Errorf("illegal length for %q: expected 5 chars but %d", w, len(w))
		}
		words = append(words, w)
	}
}

func mustLoad(b []byte) []string {
	v, err := load(b)
	if err != nil {
		panic("failed to load word list: " + err.Error())
	}
	return v
}

func init() {
	Answer = mustLoad(answer)
	Other = mustLoad(other)
}
