package wordle

type Filter interface {
	Filter(string) bool
}

type runeFilter interface {
	runeFilter([]rune) bool
}

type filters []runeFilter

func (ff filters) Filter(s string) bool {
	runes := []rune(s)
	for _, f := range ff {
		if !f.runeFilter(runes) {
			return false
		}
	}
	return true
}

type filterFunc func([]rune) bool

var falseFilter = filterFunc(func([]rune) bool { return false })

func (ff filterFunc) runeFilter(runes []rune) bool {
	return ff(runes)
}

func containsRune(runes []rune, r rune) bool {
	for _, x := range runes {
		if x == r {
			return true
		}
	}
	return false
}

func noSpotFilter(r rune) runeFilter {
	return filterFunc(func(runes []rune) bool {
		return !containsRune(runes, r)
	})
}

func correctSpotFilter(r rune, x int) runeFilter {
	if x < 0 {
		return falseFilter
	}
	return filterFunc(func(runes []rune) bool {
		return x < len(runes) && runes[x] == r
	})
}

func wrongSpotFilter(r rune, x int, ignores []int) runeFilter {
	if x < 0 {
		return falseFilter
	}
	if len(ignores) < 0 {
		return filterFunc(func(runes []rune) bool {
			return x < len(runes) && runes[x] != r && containsRune(runes, r)
		})
	}
	return filterFunc(func(runes []rune) bool {
		if x >= len(runes) || runes[x] == r {
			return false
		}
		rr := make([]rune, len(runes))
		copy(rr, runes)
		for _, n := range ignores {
			rr[n] = 0
		}
		return containsRune(rr, r)
	})
}

func ApplyFilter(f Filter, src []string) []string {
	dst := make([]string, 0, len(src) / 2)
	for _, s := range src {
		if f.Filter(s) {
			dst = append(dst, s)
		}
	}
	return dst
}
