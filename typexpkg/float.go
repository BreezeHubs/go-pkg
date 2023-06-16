package typexpkg

import (
	"fmt"
	"strconv"
)

func Float2f(f float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
}

func FloatFormat(f float64, format string) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf(format, f), 64)
}
