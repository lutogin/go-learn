package common

import (
    "testing"
)

func TestFactorial(t *testing.T) {
    t.Run("Should be equal with result", func(t *testing.T) {
        n, expectedResult := 5, 120

        realResult := Factorial(uint(n))

        if realResult != uint(expectedResult) {
            t.Errorf("Expected result %d != real result %d", expectedResult, realResult)
        }
    })
    t.Run("Should be wrong", func(t *testing.T) {
        n, expectedResult := 5, 123

        realResult := Factorial(uint(n))

        if realResult == uint(expectedResult) {
            t.Errorf("Expected result %d == real result %d", expectedResult, realResult)
        }
    })
}