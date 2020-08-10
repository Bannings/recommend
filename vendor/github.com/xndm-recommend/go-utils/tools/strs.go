package tools

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/common/slices"
	"github.com/xndm-recommend/go-utils/maths"
)

func IsEmptyStrs(s []string) bool {
	return len(s) == consts.ZERO
}

// cut strs
func CutStrList(s []string, l int) []string {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

// cut strs
func CutStrListAndFilling(s []string, f []string, l int) []string {
	if cut := CutStrList(s, l); len(cut) < l {
		tmp := append(s, f...)
		return RmDuplicateStrLen(tmp, l)
	} else {
		return cut
	}
}

// Shuffle
func ShuffleStrList(s []string) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// list自去重
func RmDuplicateStr(s []string) []string {
	if len(s) < PerfThr {
		return slices.RemoveRepByLoopStr(s)
	} else {
		return slices.RemoveRepByMapStr(s)
	}
}

func RmDuplicateStrLen(s []string, l int) []string {
	ints := RmDuplicateStr(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// s1对s2做差
func DifferenceStr(s1, s2 []string) []string {
	if len(s2) == 0 {
		return s1
	}
	ints := make([]string, 0, len(s1))
	for _, i := range s1 {
		if !slices.IsInStrSlice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func DifferenceStrLen(s1, s2 []string, l int) []string {
	ints := DifferenceStr(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// string list union
func UnionStrs(s1, s2 []string) []string {
	return append(s1, s2...)
}

func UnionStrsDup(s1, s2 []string) []string {
	return RmDuplicateStr(UnionStrs(s1, s2))
}

func UnionStrsLen(s1, s2 []string, l int) []string {
	return RmDuplicateStrLen(UnionStrs(s1, s2), l)
}

func UnionStrs32(s1, s2 []int32) []int32 {
	return append(s1, s2...)
}

// string list Loop
func GetStrListNoLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, fmt.Errorf("Input paras error!!!")
	}
	minThr := maths.MinInt(size*(num-1), len(s))
	maxThr := maths.MinInt(size*num, len(s))
	return s[minThr:maxThr], nil
}

func GetStrListLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 || size >= len(s) {
		return []string{}, fmt.Errorf("Input parameter error!!!")
	}
	start := (size * (num - 1)) % len(s)
	end := (num * size) % len(s)
	if start < end {
		return s[start:end:end], nil
	} else {
		out := make([]string, 0, size)
		out = append(out, s[start:len(s):len(s)]...)
		out = append(out, s[:end:end]...)
		return out, nil
	}
}

// ints mix
func MixListStrV2(s ...[]string) []string {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]string, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return RmDuplicateStr(mix)
}

func MixListStrLenV2(l int, s ...[]string) []string {
	mix := MixListStrV2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}
