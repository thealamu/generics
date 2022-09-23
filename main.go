package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sumInts(ints), sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))

	fmt.Printf("Generic Sums, using type constraint: %v and %v\n",
		sumNumbers(ints),
		sumNumbers(floats))
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
