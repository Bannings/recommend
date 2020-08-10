package tools

import (
	"math/rand"
	"time"

	"github.com/xndm-recommend/go-utils/maths"
)

// s1对s2做差
func DiffInterface(s1, s2 []interface{}) []interface{} {
	dup := make([]interface{}, 0, len(s1))
	for _, i := range s1 {
		sign := true
		for _, v := range s2 {
			if i == v {
				sign = false
				break
			}
		}
		if sign {
			dup = append(dup, i)
		}
	}
	return dup
}

func DiffInterfaceLen(s1, s2 []interface{}, i int) []interface{} {
	strings := DiffInterface(s1, s2)
	if i < 0 {
		return strings
	}
	return strings[:maths.MinInt(len(strings), i)]
}

func ShuffleList(s ...interface{}) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
