package shuff_test

import (
	"bytes"
	"testing"

	"github.com/kamilturek/shuff"
)

func TestShuffle(t *testing.T) {
	input := bytes.NewBufferString("hello\nworld\n")
	output := bytes.NewBufferString("")

	s, err := shuff.NewShuffler(
		shuff.WithInput(input),
		shuff.WithOutput(output),
		shuff.WithSeed(2),
	)
	if err != nil {
		t.Fatal(err)
	}

	s.Shuffle()

	want := "world\nhello\n"
	got := output.String()
	if want != got {
		t.Fatalf("want: %q, got %q", want, got)
	}
}
