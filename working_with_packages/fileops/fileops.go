package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valuetext := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuetext), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	WriteFloatToFile(0, fileName)
	if err != nil {
		return 0, errors.New("storage not found")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("storage contains invalid data")
	}
	return value, nil
}
