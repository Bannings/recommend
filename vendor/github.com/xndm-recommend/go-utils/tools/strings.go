package tools

import (
	"bytes"
	"strings"

	"github.com/xndm-recommend/go-utils/common/consts"
)

// IsEmpty returns true if the string is empty
func IsEmptyStr(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmptyStr(text string) bool {
	return !IsEmptyStr(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlankStr(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlankStr(text string) bool {
	return !IsBlankStr(text)
}

// Left justifies the text to the left
func LeftStr(text string, size int) string {
	spaces := size - len(text)
	if spaces <= 0 {
		return text
	}
	var buffer bytes.Buffer
	buffer.WriteString(text)
	for i := 0; i < spaces; i++ {
		buffer.WriteString(consts.BLANK)
	}
	return buffer.String()
}

// Right justifies the text to the right
func RightStr(text string, size int) string {
	spaces := size - len(text)
	if spaces <= 0 {
		return text
	}
	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString(consts.BLANK)
	}
	buffer.WriteString(text)
	return buffer.String()
}

// join
func JoinStr(sep string, str ...string) string {
	return strings.Join(str, sep)
}

func JoinStrSlice(sep string, str []string) string {
	return strings.Join(str, sep)
}

// buff
func JoinStrSliceByBuf(str []string) string {
	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func JoinStrByBuf(str ...string) string {
	return JoinStrSliceByBuf(str)
}

func JoinStrSliceByBufs(sep string, str []string) string {
	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
		buffer.WriteString(sep)
	}
	return buffer.String()
}

func JoinStrByBufSep(sep string, str ...string) string {
	return JoinStrSliceByBufs(sep, str)
}

// split
func SplitStrSep(str string, sep1, sep2 string) string {
	return strings.Split(strings.Split(str, sep1)[1], sep2)[0]
}

func ContainStrNum(str string, sep string) int {
	return len(strings.Split(str, sep)) - 1
}
