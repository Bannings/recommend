package maths

import (
	"github.com/pkg/errors"
)

func Min(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return MinInt(a.(int), b.(int))
		case float32:
			return MinFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	case float32:
		switch b.(type) {
		case int:
			return MinFloat(a.(float32), b.(float32))
		case float32:
			return MinFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	default:
		return errors.New("input num type error!!!")
	}
}

func Max(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return MaxInt(a.(int), b.(int))
		case float32:
			return MaxFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	case float32:
		switch b.(type) {
		case int:
			return MaxFloat(a.(float32), b.(float32))
		case float32:
			return MaxFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	default:
		return errors.New("input num type error!!!")
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxInt32(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

func MinFloat(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func MaxFloat(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}
