package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strSlice("KokoJAMBO", "эх", "страно, что вывело меня да?", ":3")
}

func strSlice(strs ...string) {
	strArr := make([]string, len(strs))
	copy(strArr, strs)

	max := ""
	for _, v := range strArr {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(max) {
			max = v
		}
	}

	fmt.Println(max)
}