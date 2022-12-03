package inputreader

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type InputReader struct{}

func (r InputReader) GetInputForDay(dayNumber int) (input string, err error) {
	filePath, err := r.GetFilePathForDay(dayNumber)
	if err != nil {
		return "", err
	}

	fileRaw, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	fileContent := string(fileRaw)

	return fileContent, nil
}

// GetFileNameForDay returns the file name for the input file in `dayxx_input.txt` format,
// where xx is the day number as a 2-digit number with leading 0's if needed
func (r InputReader) GetFileNameForDay(dayNumber int) (string, error) {
	return fmt.Sprintf("day%02v_input.txt", dayNumber), nil
}

// GetFilePathForDay returns the file path for the input file for the given day number
func (r InputReader) GetFilePathForDay(dayNumber int) (string, error) {
	fileName, err := r.GetFileNameForDay(dayNumber)
	if err != nil {
		return "", err
	}

	b, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(path.Dir(b), "..", "..", "resources", fileName)

	return filePath, nil
}
