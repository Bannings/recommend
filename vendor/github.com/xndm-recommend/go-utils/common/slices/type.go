package slices

func IntToInterface(i []int) []interface{} {
	ifs := make([]interface{}, len(i))
	for ind, v := range i {
		ifs[ind] = v
	}
	return ifs
}

func InterfaceToInt(i []interface{}) []int {
	ints := make([]int, len(i))
	for ind, v := range i {
		ints[ind] = v.(int)
	}
	return ints
}

func StrToInterface(s []string) []interface{} {
	ifs := make([]interface{}, len(s))
	for ind, v := range s {
		ifs[ind] = v
	}
	return ifs
}

func InterfaceToStr(s []interface{}) []string {
	strings := make([]string, len(s))
	for ind, v := range s {
		strings[ind] = v.(string)
	}
	return strings
}
