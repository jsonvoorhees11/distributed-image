package stringHelper

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

type CmdFlag struct {
	Name  string
	Value string
}

type argLengthError struct {
	expected int
	actual   int
}

func (e *argLengthError) Error() string {
	return "Invalid length of argument, expect " + strconv.Itoa(e.expected) + ", got " + strconv.Itoa(e.actual)
}

func GetHashCode(input string, length ...int) string {
	h := sha1.New()
	io.WriteString(h, input)
	if len(length) > 0 {
		return hex.EncodeToString(h.Sum(nil)[0:length[0]])
	} else {
		return hex.EncodeToString(h.Sum(nil))
	}
}

func ParseStrToFlag(input string) (CmdFlag, error) {
	arr := strings.Split(input, "=")
	if len(arr) < 2 || arr[0] == "" || arr[1] == "" {
		return CmdFlag{}, &argLengthError{expected: 2, actual: len(arr)}
	}
	flag := strings.TrimPrefix(arr[0], "-")
	value := arr[1]
	return CmdFlag{
		flag,
		value,
	}, nil
}
