package utils

import (
	"fmt"
	"reflect"
)

func GetTypeAsString(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func ConcatenateStrings(args ...string) string {
	var concatenatedString string
	for _, arg := range args {
		concatenatedString += fmt.Sprintf("%v", arg)
	}

	return concatenatedString
}

func StringToRunes(s string) []rune {
	return []rune(s)
}
