package goshuffle

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func shuffleWord(word string) string {

	if len(word) <= 3 {
		return word
	}

	runes := []rune(word[1 : len(word)-1])

	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return word[0:1] + string(runes) + word[len(word)-1:]
}

// New returns the shuffled string
func New(txt string) string {

	var out string

	for _, w := range strings.Split(txt, " ") {
		out += shuffleWord(w) + " "
	}
	return out[:len(out)-1]
}
