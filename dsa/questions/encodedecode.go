package main

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
