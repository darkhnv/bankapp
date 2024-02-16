package fileops

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		return 0, nil
	}
	if err != nil {
		return 0, errors.New("failed to read file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(fileName string, value float64) error {
	valueText := fmt.Sprintf("%.2f", value)
	err := ioutil.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return errors.New("failed to write to file")
	}
	return nil
}
