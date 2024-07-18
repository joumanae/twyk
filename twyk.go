package twyk

import (
	"fmt"
	"io"
	"net/http"

	"os"
	"strings"
)

// isValidURL checks if a URL is valid and properly formatted

func Match(urlString, keyword string, c *http.Client) (bool, error) {

	resp, err := c.Get(urlString)
	if resp.StatusCode != 200 {
		return false, fmt.Errorf("page not found")

	}
	if err != nil {
		return false, fmt.Errorf("an error occured %s", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}

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

	match, err := Match(args[0], args[1], http.DefaultClient)
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
