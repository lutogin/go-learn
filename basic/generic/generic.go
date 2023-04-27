package generic

import "fmt"

func IsInclude[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func RunGeneric() {
	a1 := "test"
	a2 := []string{"a", "b", "test"}
	fmt.Printf("includes a1 in a2 array: %t \n", IsInclude[string](a1, a2))

	b1 := 5
	b2 := []int{1, 2, 4, 6}

	fmt.Printf("includes b1 in b2 array: %t \n", IsInclude[int](b1, b2))
}
