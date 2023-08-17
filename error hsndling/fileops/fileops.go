package fileops

import (
	"errors"
	"os"
)

func Getfilechar(filename *string) (int, error) {
	if filename == nil {
		return 0, errors.New("nil filename")
	}
	bytes, err := os.ReadFile(*filename)
	if err != nil {
		return 0, errors.New("error in reading the file")
	}
	return len(bytes), nil
}
