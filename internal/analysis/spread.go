package analysis

import (
	"fmt"
	"strconv"
)

func CountSpread(StrbidPrice string, StraskPrice string) (float64, error) {
	ap, err := strconv.ParseFloat(StrbidPrice, 64)
	if err != nil {
		return 0, fmt.Errorf("converting string to float")
	}

	bp, err := strconv.ParseFloat(StraskPrice, 64)
	if err != nil {
		return 0, fmt.Errorf("converting string to float")
	}

	res := bp - ap
	return res, nil
}
