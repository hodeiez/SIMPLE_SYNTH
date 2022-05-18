package helpers

import "strconv"

func Float32ToString(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', 5, 32)
}
