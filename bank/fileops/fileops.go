package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read file [%s]", filename)
	}
	valueTxt := string(fileBytes)
	valueFloat, err := strconv.ParseFloat(valueTxt, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse float value from string [%s]", valueTxt)
	}
	return valueFloat, nil
}

func WriteFloatToFile(filename string, val float64) {
	valueTxt := fmt.Sprint(val)
	os.WriteFile(filename, []byte(valueTxt), 0644)
}
