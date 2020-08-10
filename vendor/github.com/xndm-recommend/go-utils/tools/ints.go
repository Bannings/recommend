package tools

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/common/slices"
	"github.com/xndm-recommend/go-utils/maths"
)

var PerfThr = 300

func IsEmptyInts(s []int) bool {
	return len(s) == consts.ZERO
}

func IsEmptyInt32s(s []int32) bool {
	return len(s) == consts.ZERO
}

// cut strs
func CutIntList(s []int, l int) []int {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

// cut strs
func CutIntListAndFilling(s []int, f []int, l int) []int {
	if cut := CutIntList(s, l); len(cut) < l {
		tmp := append(s, f...)
		return RmDuplicateIntLen(tmp, l)
	} else {
		return cut
	}
}

// cut strs
func CutInt32List(s []int32, l int) []int32 {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

// cut strs
func CutInt32ListAndFilling(s []int32, f []int32, l int) []int32 {
	if cut := CutInt32List(s, l); len(cut) < l {
		tmp := append(s, f...)
		return RmDuplicateInt32Len(tmp, l)
	} else {
		return cut
	}
}

// Shuffle
func ShuffleIntList(s []int) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func ShuffleInt32List(s []int32) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// list自去重
func RmDuplicateInt(s []int) []int {
	if len(s) < PerfThr {
		return slices.RemoveRepByLoop(s)
	} else {
		return slices.RemoveRepByMap(s)
	}
}

func RmDuplicateInt32(s []int32) []int32 {
	if len(s) < PerfThr {
		return slices.RemoveRepByLoop32(s)
	} else {
		return slices.RemoveRepByMap32(s)
	}
}

func RmDuplicateIntLen(s []int, l int) []int {
	ints := RmDuplicateInt(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func RmDuplicateInt32Len(s []int32, l int) []int32 {
	ints := RmDuplicateInt32(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// s1对s2做差
func DifferenceInt(s1, s2 []int) []int {
	if len(s2) == 0 {
		return s1
	}
	ints := make([]int, 0, len(s1))
	for _, i := range s1 {
		if !slices.IsInIntSlice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func DifferenceIntLen(s1, s2 []int, l int) []int {
	ints := DifferenceInt(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func DifferenceInt32(s1, s2 []int32) []int32 {
	if len(s2) == 0 {
		return s1
	}
	ints := make([]int32, 0, len(s1))
	for _, i := range s1 {
		if !slices.IsInInt32Slice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func DifferenceInt32Len(s1, s2 []int32, l int) []int32 {
	ints := DifferenceInt32(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// int list union
func UnionInts(s1, s2 []int) []int {
	return append(s1, s2...)
}

func UnionIntsDup(s1, s2 []int) []int {
	return RmDuplicateInt(UnionInts(s1, s2))
}

func UnionIntsLen(s1, s2 []int, l int) []int {
	return RmDuplicateIntLen(UnionInts(s1, s2), l)
}

func UnionInts32(s1, s2 []int32) []int32 {
	return append(s1, s2...)
}

func UnionInts32Dup(s1, s2 []int32) []int32 {
	return RmDuplicateInt32(UnionInts32(s1, s2))
}

func UnionInts32Len(s1, s2 []int32, l int) []int32 {
	return RmDuplicateInt32Len(UnionInts32(s1, s2), l)
}

// int list Loop
func GetIntListNoLoop(s []int, size, num int) ([]int, error) {
	if num <= 0 || size <= 0 {
		return []int{}, fmt.Errorf("Input paras error!!!")
	}
	minThr := maths.MinInt(size*(num-1), len(s))
	maxThr := maths.MinInt(size*num, len(s))
	return s[minThr:maxThr], nil
}

func GetIntListLoop(s []int, size, num int) ([]int, error) {
	if num <= 0 || size <= 0 || size >= len(s) {
		return []int{}, fmt.Errorf("Input parameter error!!!")
	}
	start := (size * (num - 1)) % len(s)
	end := (num * size) % len(s)
	if start < end {
		return s[start:end:end], nil
	} else {
		out := make([]int, 0, size)
		out = append(out, s[start:len(s):len(s)]...)
		out = append(out, s[:end:end]...)
		return out, nil
	}
}

// ints mix
func MixListIntV2(s ...[]int) []int {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]int, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return RmDuplicateInt(mix)
}

func MixListIntLenV2(l int, s ...[]int) []int {
	mix := MixListIntV2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}

func MixListInt32V2(s ...[]int32) []int32 {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]int32, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return RmDuplicateInt32(mix)
}

func MixListInt32LenV2(l int, s ...[]int32) []int32 {
	mix := MixListInt32V2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}
