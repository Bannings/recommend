package slices

import "math/rand"

func GenRandInt(size int) []int {
	ints := make([]int, size)
	for ind := range ints {
		ints[ind] = rand.Int()
	}
	return ints
}

func GenNRandInt(size int, n int) []int {
	ints := make([]int, size)
	for ind := range ints {
		ints[ind] = rand.Intn(n)
	}
	return ints
}

func GenRandInt32(size int) []int32 {
	ints := make([]int32, size)
	for ind := range ints {
		ints[ind] = rand.Int31()
	}
	return ints
}

func GenNRandInt32(size int, n int32) []int32 {
	ints := make([]int32, size)
	for ind := range ints {
		ints[ind] = rand.Int31n(n)
	}
	return ints
}
