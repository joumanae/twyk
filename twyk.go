package twyk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Matcher struct {
	HTTPClient *http.Client
}

func NewMatcher() *Matcher {
	return &Matcher{
		HTTPClient: http.DefaultClient,
	}
}

func (m *Matcher) Match(uri, keyword string) (bool, error) {
	resp, err := m.HTTPClient.Get(uri)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("page not found: %s", uri)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if bytes.Contains(body, []byte(keyword)) {
		return true, nil
	}
	return false, nil
}

func Match(uri, keyword string) (bool, error) {
	m := NewMatcher()
	return m.Match(uri, keyword)
}

func Main() int {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: twyk <url> <keyword>")
		os.Exit(0)
	}
	uri, keyword := args[0], args[1]
	matcher := NewMatcher()
	match, err := matcher.Match(uri, keyword)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s: ", keyword)
	if match {
		fmt.Println("match")
	} else {
		fmt.Println("no match")
	}

	return 0
}
