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
