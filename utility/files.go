package utility

import (
	"errors"
	"os"
)

func ReadFromFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("can't read data from file")
	}

	return string(data), nil
}

func WriteToFile(fileName, data string) error {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return errors.New("can't write data to file")
	}

	return nil
}
