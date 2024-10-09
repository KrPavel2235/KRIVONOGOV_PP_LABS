package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

}

