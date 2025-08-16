package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatVal, error := strconv.ParseFloat(stringVal, 64)
		if error != nil {
			fmt.Println(error)
			return nil, errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
