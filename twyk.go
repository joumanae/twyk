package twyk

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Matcher struct {
	Out        io.Writer
	HTTPClient *http.Client
}

func NewMatcher() *Matcher {
	return &Matcher{
		Out:        os.Stdout,
		HTTPClient: http.DefaultClient,
	}
}

func (m *Matcher) PrintMatch(uri, keyword string) error {
	matched, err := m.Match(uri, keyword)
	if err != nil {
		return err
	}
	fmt.Fprintf(m.Out, "%s: ", keyword)
	if matched {
		fmt.Fprintln(m.Out, "match")
	} else {
		fmt.Fprintln(m.Out, "no match")
	}
	return nil
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
	return bytes.Contains(body, []byte(keyword)), nil
}

func PrintMatch(uri, keyword string) error {
	m := NewMatcher()
	return m.PrintMatch(uri, keyword)
}

func Main() int {
	path := flag.String("f", "", "path to file of URLs and keywords")
	flag.Parse()
	if *path != "" {
		err := CheckURLsFromFile(*path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if flag.NArg() != 2 {
		fmt.Println("Usage: twyk <url> <keyword>")
		return 0
	}
	uri, keyword := flag.Args()[0], flag.Args()[1]
	err := PrintMatch(uri, keyword)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func CheckURLsFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		res := strings.Split(input.Text(), " ")
		if len(res) != 2 {
			continue
		}
		uri, keyword := res[0], res[1]
		err := PrintMatch(uri, keyword)
		if err != nil {
			return err
		}
	}
	err = input.Err()
	if err != nil {
		return err
	}
	return nil
}
