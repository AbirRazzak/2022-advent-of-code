package inputreader

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InputReader interface {
	GetInputForDay(dayNumber int) (string, error)
}

type HTTPInputReader struct{}

func (r *HTTPInputReader) GetInputForDay(dayNumber int) (string, error) {
	url, err := r.GetURLPathForFileName(dayNumber)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// remove the trailing new line at the end of the input string
	input := strings.Trim(string(body), "\n")

	return input, nil
}

func (r *HTTPInputReader) GetURLPathForFileName(dayNumber int) (string, error) {
	return fmt.Sprintf("https://raw.githubusercontent.com/AbirRazzak/2022-advent-of-code/master/resources/day%02v_input.txt", dayNumber), nil
}
