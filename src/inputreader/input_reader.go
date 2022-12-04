package inputreader

import (
	"fmt"
	"io"
	"net/http"
)

type InputReader interface {
	GetInputForDay(dayNumber int) (input string, err error)
}

type HTTPInputReader struct{}

func (r *HTTPInputReader) GetInputForDay(dayNumber int) (input string, err error) {
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

	return string(body), nil
}

func (r *HTTPInputReader) GetURLPathForFileName(dayNumber int) (string, error) {
	return fmt.Sprintf("https://raw.githubusercontent.com/AbirRazzak/2022-advent-of-code/master/resources/day%02v_input.txt", dayNumber), nil
}
