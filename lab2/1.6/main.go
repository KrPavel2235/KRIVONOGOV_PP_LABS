package main

import (
	"fmt"
)

func main() {
	//Задание 6
	fmt.Println("AVG чисел 6 и 8")
	AVG(6,8)
	
	
}	
	func AVG(num...int){
		sum := 0
		count := 0
		res := 0.0
		for _, v := range num{
			count++
			sum += v
			
		}
		res = float64(sum/count)
		fmt.Println(res)
	}