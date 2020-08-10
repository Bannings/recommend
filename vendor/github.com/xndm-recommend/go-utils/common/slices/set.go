package slices

// int slice自去重,通过两重循环过滤重复元素
func RemoveRepByLoop(s []int) []int {
	dup := make([]int, 0, len(s))
	for _, v := range s {
		if !IsInIntSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(s []int) []int {
	dup := make([]int, 0, len(s))
	tmpMap := make(map[int]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func RemoveRepByLoop32(s []int32) []int32 {
	dup := make([]int32, 0, len(s))
	for _, v := range s {
		if !IsInInt32Slice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap32(s []int32) []int32 {
	dup := make([]int32, 0, len(s))
	tmpMap := make(map[int32]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func RemoveRepByLoopStr(s []string) []string {
	dup := make([]string, 0, len(s))
	for _, v := range s {
		if !IsInStrSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMapStr(s []string) []string {
	dup := make([]string, 0, len(s))
	tmpMap := make(map[string]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}
