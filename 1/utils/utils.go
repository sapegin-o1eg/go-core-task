package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func GetVarType(args ...interface{}) {
	for index, arg := range args {
		fmt.Printf("Argument #%d has value: %v and type: %T\n", index+1, arg, arg)
	}
}

func HashRunesWithSalt(runes []rune, salt ...string) string {
	inputString := string(runes)

	mid := len(inputString) / 2
	var saltStr string
	if len(salt) == 0 {
		saltStr = "go-2024"
	} else {
		saltStr = salt[0]
	}

	saltedString := inputString[:mid] + saltStr + inputString[mid:]

	hasher := sha256.New()
	hasher.Write([]byte(saltedString))
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}
