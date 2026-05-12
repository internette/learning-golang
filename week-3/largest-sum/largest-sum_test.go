package main

import (
	"reflect"
	"testing"
)

func TestConvertFlagStringToIntArr(t *testing.T) {
	input := "1,2,3"
	flags := convertFlagStringToIntArr(input)
	expectedValue := []int{1, 2, 3}
	if !reflect.DeepEqual(expectedValue, flags) {
		t.Errorf("Expected %v, got %v", expectedValue, flags)
	}
}

func TestSumArray(t *testing.T) {
	input := []int{1, 2, -4, 3, -11, 8, 2}
	result := sumArray(input)
	expectedValue := 1
	if result != expectedValue {
		t.Errorf("Expected %v, got %v", expectedValue, result)
	}
}

func TestGetLargestSumAndArray(t *testing.T) {
	input := []int{1, 2, -4, 3, -11, 8, 2}
	largestSum, largestArray := getLargestSumAndArray(input)
	expectedSum := 10
	expectedArray := []int{8, 2}
	if largestSum != expectedSum {
		t.Errorf("Expected sum %v, got %v", expectedSum, largestSum)
	}
	if !reflect.DeepEqual(expectedArray, largestArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, largestArray)
	}
}
