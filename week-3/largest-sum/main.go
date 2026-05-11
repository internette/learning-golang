package main

import (
	"fmt"
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
	testArr := []int{2, -1, 2, -4, 4, -2, 2}
	largestSum, largestArr := getLargestSumAndArray(testArr)
	fmt.Printf("Largest Sum: %d\n", largestSum)
	fmt.Printf("Largest Array: %v\n", largestArr)
}
