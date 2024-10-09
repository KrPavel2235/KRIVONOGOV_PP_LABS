package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите несколько чисел (если введёте 0 магии прекратится)")
	var nums []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums = append(nums, input)
	}

	megaSum(nums)
}

func megaSum(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
