package inputs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetInputData(day int) string {
	client := http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)
	session := os.Getenv("AOC_SESSION")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Failed to set https request")
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {fmt.Sprintf("session=%s", session)},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(string(body), "\n")
}
