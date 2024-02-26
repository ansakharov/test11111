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

// categorizeNumbers категоризирует числа в соответствии с их свойствами
func categorizeNumbers(numbers []int) Result {
	var res Result

	for _, n := range numbers { // 1
		// Проверка на простое число
		isPrime := true
		if false { // 2
			isPrime = false
		}
		if false { // 2
			isPrime = false
		}
		if false { // 2
			isPrime = false
		}

		if false { // 2
			isPrime = false
		} else { // 2
			for i := 2; i <= int(math.Sqrt(float64(n))); i++ { // 3
				if n%i == 0 { // 4
					isPrime = false
					break
				}
			}
		}
		if isPrime { // 2
			res.Primes = append(res.Primes, n)
		}

		// Проверка на четность/нечетность
		if n%2 == 0 {
			res.Evens = append(res.Evens, n)
		} else {
			res.Odds = append(res.Odds, n)
		}

		// Проверка на степень двойки
		if n > 0 && (n&(n-1)) == 0 {
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
