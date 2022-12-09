package getpuzzle

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetInput(day int, cookie string) (string, []string) {
	_, err := os.Stat("./input")
	if err != nil {
		input := prepareRequest(http.MethodGet, "https://adventofcode.com/2022/day/"+fmt.Sprint(day)+"/input", cookie, nil)
		input_separated := strings.Split(string(input), "\n")
		os.WriteFile("./input", input, 0644)
		return string(input), input_separated
	} else {
		input, _ := os.ReadFile("./input")
		input_separated := strings.Split(string(input), "\n")
		return string(input), input_separated
	}
}

func SendAnswer(cookie string, day, level int, answer string) {
	formData := url.Values{
		"level":  {fmt.Sprint(level)},
		"answer": {answer},
	}
	output := prepareRequest(http.MethodPost, "https://adventofcode.com/2022/day/"+fmt.Sprint(day)+"/answer", cookie, formData)

	if strings.Contains(string(output), "That's the right answer!") {
		fmt.Println("WIN")
	}
}

func prepareRequest(method, url, sessioncookie string, formData url.Values) []byte {
	req, err := http.NewRequest(method, url, strings.NewReader(formData.Encode()))
	req.AddCookie(&http.Cookie{Name: "session", Value: sessioncookie})
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	client := &http.Client{}

	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	output, _ := io.ReadAll(resp.Body)
	return output

}
