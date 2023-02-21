package common

func Add(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SimpleAdd(a int, b int) int {
	return a + b
}

func FiboTest(init, final int) (int, int) {
	a := init + 1
	b := final + 2
	return a, b
}

func Action(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func Square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func Fibo(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibo(n-1) + Fibo(n-2)
}

func Devide(a uint, b uint) float32 {
	if b < 1 {
		panic("Devide by zero impossible!")
	}

	return float32(a) / float32(b)
}
