package utils

import "fmt"

func ConvertAndCatToString(args ...interface{}) string {
	var result string
	for _, arg := range args {
		result += fmt.Sprint(arg)
	}

	return result
}