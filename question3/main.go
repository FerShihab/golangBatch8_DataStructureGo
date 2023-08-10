package main

import (
	"fmt"
)

func munculSekali(inputString string) []int {
	numCount := make(map[int]int)

	for _, char := range inputString {
		num := int(char - '0')
		numCount[num]++
	}

	var uniqueNumbers []int
	for num, count := range numCount {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	return uniqueNumbers
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
