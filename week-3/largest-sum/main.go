package main

import (
	"flag"
	"fmt"
	"regexp"
)

func sumArray(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func getLargestSumAndArray(numbers []int) (int, []int) {
	largestSum := sumArray(numbers)
	largestArr := numbers
	if len(numbers) == 0 {
		return 0, nil
	}
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers)+1; j++ {
			currentArr := numbers[i:j]
			currentSum := sumArray(currentArr)
			fmt.Println(currentArr, currentSum)
			if len(currentArr) > 1 && currentSum > largestSum {
				largestSum = currentSum
				largestArr = currentArr
			}
		}
	}
	return largestSum, largestArr
}

func main() {
	input := flag.String("input", "", "Input array of integers")
	flag.Parse()
	matched, err := regexp.MatchString(`[a-zA-Z]+`, *input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if matched {
		fmt.Println("Please provide a valid input array of two or more integers")
		return
	}
	// largestSum, largestArr := getLargestSumAndArray(input)
	// fmt.Printf("Largest Sum: %d\n", largestSum)
	// fmt.Printf("Largest Array: %v\n", largestArr)
}
