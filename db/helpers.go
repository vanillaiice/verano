package db

import (
	"fmt"
	"strconv"
	"strings"
)

// flat converts a slice of integers into a comma-separated string.
func flat(vals []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vals)), ","), "[]")
}

// unflat converts a comma-separated string into a slice of integers.
func unflat(s string) ([]int, error) {
	var intVals []int
	if s == "" {
		return intVals, nil
	}
	sVals := strings.Split(s, ",")
	for _, s := range sVals {
		val, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		intVals = append(intVals, val)
	}
	return intVals, nil
}
