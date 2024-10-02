package stringutils

import "fmt"

func ReverseString(str string) {
	res := ""

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	fmt.Println(res)
}