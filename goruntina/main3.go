package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	numbers := make(chan int)
	sum := 0

	go func() {
		for i := 1; i <= 10; i++ {
			numbers <- i * i
		}
		close(numbers)
	}()

	for n := range numbers {
		sum += n
	}

	fmt.Println("Sum of squares from 1 to 10:", sum)
	//=========================================================
	//11
	result := make(chan int)

	go sumOfSquares(1, 1000, result)

	fmt.Println("Sum of squares from 1 to 1000:", <-result)

	//12 ================================================

	text := "Hello World"
	result := make(chan map[rune]int)

	go countVowelsConsonants(text, result)

	counts := <-result

	fmt.Println("Vowels:", counts['v'])
	fmt.Println("Consonants:", counts['c'])
}

func sumOfSquares(start, end int, result chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i * i
	}
	result <- sum

}

func countVowelsConsonants(text string, result chan map[rune]int) {
	counts := make(map[rune]int)

	for _, char := range text {
		if strings.ContainsRune("aeiouAEIOU", char) {
			counts['v']++
		} else if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			counts['c']++
		}
	}

	result <- counts

}

//==========================

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(start, end int, result chan int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			result <- i
		}
	}
	close(result)
}

func main() {
	result := make(chan int)

	go findPrimes(1, 100, result)

	for prime := range result {
		fmt.Println(prime)
	}
}

func countWords(text string, result chan int) {
	words := strings.Fields(text)
	result <- len(words)
}

func main() {
	text := "This is a sample text for counting words concurrently"
	result := make(chan int)

	go countWords(text, result)

	fmt.Println("Number of words:", <-result)
}

func computeSquareRoots(numbers []float64, result chan float64) {
	for _, num := range numbers {
		result <- math.Sqrt(num)
	}
	close(result)
}

func main() {
	numbers := []float64{16, 25, 36, 49, 64}
	result := make(chan float64)

	go computeSquareRoots(numbers, result)

	for squareRoot := range result {
		fmt.Println(squareRoot)
	}
}

func fibonacci(n int, result chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		result <- a
		a, b = b, a+b
	}
	close(result)
}

func main() {
	n := 10
	result := make(chan int)

	go fibonacci(n, result)

	for num := range result {
		fmt.Println(num)
	}
}

func sumEven(numbers []int, result chan int) {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	result <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int)

	go sumEven(numbers, result)

	fmt.Println("Sum of even numbers:", <-result)
}

func findLongestWordLength(text string, result chan int) {
	words := strings.Fields(text)
	longestLength := 0
	for _, word := range words {
		if len(word) > longestLength {
			longestLength = len(word)
		}
	}
	result <- longestLength
}

func main() {
	text := "This is a sample text to find the longest word length concurrently"
	result := make(chan int)

	go findLongestWordLength(text, result)

	fmt.Println("Length of the longest word:", <-result)
}
