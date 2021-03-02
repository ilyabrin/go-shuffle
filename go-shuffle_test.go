package goshuffle

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(0)
}

func TestNew(t *testing.T) {

	tests := []struct {
		got  string
		want string
	}{
		{New("Random"), "Rdnaom"},
		{New("No"), "No"},
		{New("Yes"), "Yes"},
		// TODO: add more cases
	}

	for _, tc := range tests {
		if tc.got != tc.want {
			t.Errorf("got: %v but want %v", tc.got, tc.want)
		}
	}

}
