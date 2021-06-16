package fibonacci

import "fmt"

var memo = map[int]int{}

func FindFibo(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] != 0 {
		fmt.Println(memo[n])
		return int(memo[n])
	}
	a := FindFibo(n - 1)
	b := FindFibo(n - 2)
	memo[n-1] = a
	memo[n-2] = b
	return a + b
}
