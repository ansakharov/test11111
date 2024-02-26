package main

import (
	"fmt"
	"math"
)

// Result содержит различные категории чисел
type Result struct {
	Primes    []int
	Evens     []int
	Odds      []int
	PowersOf2 []int
}

// isPrime проверяет, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPowerOfTwo проверяет, является ли число степенью двойки
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// categorizeNumbers категоризирует числа в соответствии с их свойствами
func categorizeNumbers(numbers []int) Result {
	var res Result

	for _, n := range numbers {
		if isPrime(n) {
			res.Primes = append(res.Primes, n)
		}
		if n%2 == 0 {
			res.Evens = append(res.Evens, n)
		} else {
			res.Odds = append(res.Odds, n)
		}
		if isPowerOfTwo(n) {
			res.PowersOf2 = append(res.PowersOf2, n)
		}
	}

	return res
}

func main() {
	// Создаем массив чисел от 0 до 1024
	var numbers []int
	for i := 0; i <= 1024; i++ {
		numbers = append(numbers, i)
	}

	// Получаем категоризированные числа
	categorizedNumbers := categorizeNumbers(numbers)

	fmt.Println("Primes:", categorizedNumbers.Primes)
	fmt.Println("Evens:", categorizedNumbers.Evens)
	fmt.Println("Odds:", categorizedNumbers.Odds)
	fmt.Println("Powers of 2:", categorizedNumbers.PowersOf2)
}
