package twyk_test

import (
	"testing"

	"github.com/joumanae/twyk"
)

func TestMatch(t *testing.T) {
	got, err := twyk.Match("https://github.com/joumanae/twyk", "joumanae")
	if err != nil {
		t.Error(err)
	}
	want := true
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
