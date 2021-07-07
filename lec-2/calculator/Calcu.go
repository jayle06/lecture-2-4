package calculator

import "fmt"

func FindMin(arr []int) int {
	min := arr[0]
	for _, item := range arr {
		if item < min {
			min = item
		}
	}
	return min
}

func FindMax(arr []int) int {
	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}

func Average(arr []int) (result float64) {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	if len(arr) == 0 {
		return 0
	} else {
		result = float64(sum) / float64(len(arr))
		return result
	}
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func CheckArrPrime(arr []int) {
	for _, item := range arr {
		if IsPrime(item) {
			fmt.Printf("%d is prime number\n", item)
		} else {
			fmt.Printf("%d isn't prime number\n", item)
		}
	}
}
