// Package jellyfish provides phonetic encoding and string comparison functions.
// It is a Go port of the Python jellyfish library (https://github.com/jamesturk/jellyfish).
package jellyfish

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func normalize(str string) []rune {
	str = string(norm.NFKD.Bytes([]byte(str)))
	return []rune(strings.ToUpper(str))
}
