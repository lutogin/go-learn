package common

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	/*
	   This func will be run all time, even there will be panic inside tests, should be initialized up in the test
	*/
	t.Cleanup(func() {
		fmt.Println("Clean all func")
	})

	t.Run("Should be equal with result", func(t *testing.T) {
		n, expectedResult := 5, 120

		realResult := Factorial(uint(n))

		if realResult != uint(expectedResult) {
			t.Errorf("Expected result %d != real result %d", expectedResult, realResult)
		}
	})
	t.Run("Should be wron g", func(t *testing.T) {
		n, expectedResult := 5, 123

		realResult := Factorial(uint(n))

		if realResult == uint(expectedResult) {
			t.Errorf("Expected result %d == real result %d", expectedResult, realResult)
		}
	})

}
