package goutils

import (
	"errors"
	"math"
	"sort"

	mapset "github.com/deckarep/golang-set"
)

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

// VarianceFloat64 求方差
func VarianceFloat64(fs []float64) (float64, error) {
	// 均值
	favg, err := AvgFloat64(fs)
	if err != nil {
		return 0, err
	}
	variance := float64(0)
	for _, f := range fs {
		variance += math.Pow(f-favg, 2)
	}
	fsLen := len(fs)
	if fsLen < 2 {
		return variance, nil
	}
	variance = variance / (float64(len(fs) - 1))
	return variance, nil
}

// StdDeviationFloat64 求标准差
func StdDeviationFloat64(fs []float64) (float64, error) {
	v, err := VarianceFloat64(fs)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(v), nil
}

// MidValueFloat64 获取中位数
func MidValueFloat64(values []float64) (float64, error) {
	vlen := len(values)
	if vlen == 0 {
		return 0, errors.New("no data")
	}
	sort.Float64s(values)
	mid := vlen / 2
	if vlen%2 == 0 {
		return (values[mid-1] + values[mid]) / 2.0, nil
	}
	return values[mid], nil

}

// JaccardSimilarity Jaccard similarity相似性系数
// J(A, B) = |A ∩ B | / | A ∪ B |
// A,B分别代表符合某种条件的集合：两个集合交集的大小/两个集合并集的大小，交集=并集意味着2个集合完全重合。
func JaccardSimilarity(s1, s2 []interface{}) float64 {
	A := mapset.NewSetFromSlice(s1)
	B := mapset.NewSetFromSlice(s2)
	AiB := A.Intersect(B)
	AuB := A.Union(B)
	return float64(AiB.Cardinality()) / float64(AuB.Cardinality())
}
