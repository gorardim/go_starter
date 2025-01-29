package utils

import "unicode"

type WordScanner struct {
	s     []rune
	start int
	end   int
}

// NewWordScanner returns a new WordScanner
func NewWordScanner(s string) *WordScanner {
	return &WordScanner{
		s: []rune(s),
	}
}

func (w *WordScanner) NextWord() string {

	for w.end < len(w.s) {
		r := w.s[w.end]
		if unicode.IsSpace(r) {
			if w.start == w.end {
				w.start++
				w.end++
				continue
			}
			value := string(w.s[w.start:w.end])
			w.end++
			w.start = w.end
			return value
		}
		w.end++
	}

	return string(w.s[w.start:w.end])
}

// Rest word
func (w *WordScanner) Rest() string {
	if w.end == len(w.s) {
		return ""
	}
	return string(w.s[w.end:])
}
