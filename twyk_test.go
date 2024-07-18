package twyk_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joumanae/twyk"
)

func TestMatch(t *testing.T) {
	got, err := twyk.Match("https://github.com/joumanae/twyk", "joumanae", http.DefaultClient)
	if err != nil {
		t.Error(err)
	}
	want := true
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestMatchReturnsErrorForInvalidUrl(t *testing.T) {
	s := httptest.NewTLSServer(nil)
	defer s.Close()
	c := s.Client()
	r, err := c.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	if r.StatusCode != 404 {
		t.Errorf("wrong status code: %d", r.StatusCode)
	}
	_, err = twyk.Match(s.URL, "joumanae", c)
	if err == nil {
		t.Errorf("failed")
	}
	r.Body.Close()
}
