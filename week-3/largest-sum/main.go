package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertFlagStringToIntArr(input string) []int {
	slice := strings.Split(input, ",")
	intArr := make([]int, len(slice))
	for i, v := range slice {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		intArr[i] = int(f)
	}
	return intArr
}

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
	parsedFlagInput := convertFlagStringToIntArr(*input)
	largestSum, largestArr := getLargestSumAndArray(parsedFlagInput)
	fmt.Printf("Largest Sum: %d\n", largestSum)
	fmt.Printf("Largest Array: %v\n", largestArr)
}
