package twyk

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Match(url, keyword string) (bool, error) {
	// upload the page from the url
	fmt.Println("fetching", url)
	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	// see if there is a match of the keyword on that page
	sr := string(b)
	if strings.Contains(sr, keyword) {
		return true, nil
	}
	return false, nil
}

func Main() int {

	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("Usage: twyk <url> <keyword>")
		os.Exit(1)
	}

	match, err := Match(args[0], args[1])
	if err != nil {
		fmt.Println("Error matching keyword:", err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		os.Exit(1)
	}

	if match {
		fmt.Println("match")
	} else {
		fmt.Println("no match")
	}
	return 0
}
