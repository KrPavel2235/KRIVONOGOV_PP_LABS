package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите несколько чисел (Если введёте 0 волшебник обидется)")
	var nums2 []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums2 = append(nums2, input)
	}
	reverseArrOfNums(nums2)
}


func reverseArrOfNums(nums []int) {
	for _, v := range nums {
		defer fmt.Print(v," ")
	}
}
