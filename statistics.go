package goutils

import "errors"

// AvgFloat64 平均值
func AvgFloat64(f []float64) (float64, error) {
	fl := len(f)
	if fl == 0 {
		return 0, errors.New("empty slice")
	}
	sum := float64(0)
	for _, i := range f {
		sum += i
	}
	return sum / float64(len(f)), nil
}
