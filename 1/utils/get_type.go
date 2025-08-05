package utils

import (
	"fmt"
)

func GetVarType(args ...interface{}) {
	for index, arg := range args {
		fmt.Printf("Argument #%d has value: %v and type: %T\n", index+1, arg, arg)
	}
}
