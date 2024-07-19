package twyk_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joumanae/twyk"
)

func TestMatchReturnsTrueWhenKeywordIsMatched(t *testing.T) {
	s := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello from the test handler!")
	}))
	m := twyk.NewMatcher()
	m.HTTPClient = s.Client()
	matched, err := m.Match(s.URL, "hello")
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("expected to match %q, but did not", "hello")
	}
}

func TestMatchReturnsErrorForNotFoundUrl(t *testing.T) {
	s := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer s.Close()
	m := twyk.NewMatcher()
	m.HTTPClient = s.Client()
	_, err := m.Match(s.URL, "joumanae")
	if err == nil {
		t.Error("no error")
	}
}
