package main

/*
	1. Encode like: len(str) + "#" + str -> 6#lokesh
		- # -> Separator
		- len(str) -> After separator how much character we have to take
	2. Decode -> Iterate over the encoded string
		- If char == '#'
			- wordLen := int(encodedStr[ind-1]) - '0'
			- word := encodedStr[ind+1 : ind + wordLen + 1]
			- res = append(res, word)
			- ind = ind + wordLength
		- Otherwise increments
*/

import (
	"strconv"
)

func encode(strList []string) (encodedStr string) {
	for _, str := range strList {
		encodedStr += strconv.Itoa(len(str)) + "#" + str
	}
	return encodedStr
}

func decode(encodedStr string) (strList []string) {
	for ind := 0; ind < len(encodedStr); ind++ {
		if encodedStr[ind] == '#' {
			wordLength := int(encodedStr[ind-1] - '0')
			word := encodedStr[ind+1 : ind+wordLength+1]
			strList = append(strList, word)
			ind = ind + wordLength
		}
	}
	return strList
}
