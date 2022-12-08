package inputreader_test

import (
	"github.com/AbirRazzak/2022-advent-of-code/inputreader"
	"testing"
)

func TestHTTPInputReader_GetInputForDay(t *testing.T) {
	t.Run("get test input file", func(t *testing.T) {
		r := &inputreader.HTTPInputReader{}
		want := "this is a test file"

		got, err := r.GetInputForDay(0)
		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("GetInputForDay() got = %v, want %v", got, want)
		}
	})
}
