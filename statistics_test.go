package goutils

import (
	"fmt"
	"testing"
)

func TestJaccardSimilarity(t *testing.T) {
	s1 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s2 := []string{"0", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	s3 := []string{"0", "1", "12", "13", "14", "15", "16", "17", "18", "19"}
	s4 := []string{"0", "1", "2", "13", "14", "15", "16", "17", "18", "19"}
	s5 := []string{"0", "1", "2", "3", "14", "15", "16", "17", "18", "19"}

	all := [][]interface{}{}

	si1 := []interface{}{}
	for _, i := range s1 {
		si1 = append(si1, i)
	}
	all = append(all, si1)
	si2 := []interface{}{}
	for _, i := range s2 {
		si2 = append(si2, i)
	}
	all = append(all, si2)
	si3 := []interface{}{}
	for _, i := range s3 {
		si3 = append(si3, i)
	}
	all = append(all, si3)
	si4 := []interface{}{}
	for _, i := range s4 {
		si4 = append(si4, i)
	}
	all = append(all, si4)
	si5 := []interface{}{}
	for _, i := range s5 {
		si5 = append(si5, i)
	}
	all = append(all, si5)

	for i := 0; i < len(all); i++ {
		for j := i + 1; j < len(all); j++ {
			fmt.Printf("si:%v sj:%v sim:%v\n", all[i], all[j], JaccardSimilarity(all[i], all[j]))
		}
	}
}
