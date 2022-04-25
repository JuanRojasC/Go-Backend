package fibonacci

// RECURSIVE METHOD

func CalculateFibonacciRecursive(value int) int {
	if value < 2 {
		return value
	}
	return CalculateFibonacciRecursive(value-1) + CalculateFibonacciRecursive(value-2)
}

func GenerateFibonacciRecursive(value int) []int {
	fibo := []int{}
	for i := 0; len(fibo) < value+1; i++ {
		fibo = append(fibo, CalculateFibonacciRecursive(i))
	}
	return fibo
}

// ITERATIVE METHOD
func CalculateFibonacciIterative(value int) int {
	if value < 2 {
		return value
	}
	n, c, t := 1, 0, 0
	for i := 1; i <= value; i++ {
		t = c
		c = n
		n = n + t
	}
	return c
}

func GenerateFibonacciIterative(value int) []int {
	fibo := []int{0, 1}
	for i := 2; i <= value; i++ {
		fibo = append(fibo, fibo[i-2]+fibo[i-1])
	}
	return fibo
}

// func main() {
// 	rv := 6
// 	fmt.Println(CalculateFibonacciRecursive(rv))
// 	fmt.Println(GenerateFibonacciRecursive(rv))
// }
