package slices

import "github.com/xndm-recommend/go-utils/common/consts"

func IsInIntSlice(s []int, item int) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsInInt32Slice(s []int32, item int32) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsInStrSlice(s []string, item string) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}
