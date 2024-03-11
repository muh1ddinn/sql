package main

import (
	"fmt"
	"sort"
)

func ex1() {

	ch := make(chan int)

	go func() {

		for i := 1; i <= 10; i++ {

			ch <- i

		}
		close(ch)

	}()

	for n := range ch {
		fmt.Println(n)
	}

}

func ex2() {

	ch := make(chan int)

	go func() {

		sum := 0
		for i := 0; i < 100; i++ {

			sum = sum + i
			ch <- sum
		}
		close(ch)

	}()

	for n := range ch {

		fmt.Println(n)

	}

}

func ex3(l int) {

	ch := make(chan int)

	sum := 1

	if l == 0 {
		fmt.Println(l)
	} else if l > 0 {

		go func() {

			for i := 1; i < l; i++ {

				sum = sum * i
				ch <- sum
			}
			close(ch)

		}()

		for n := range ch {

			fmt.Println(n)

		}
	}
}

func ex4(p int) {

	ch := make(chan int)

	go func() {
		defer close(ch)

		sum := 0
		if p%2 == 0 {
			for i := 0; i <= p; i += 2 {
				sum += i
				ch <- sum
			}
		} else if p%2 != 0 {
			p++

			for i := 0; i <= p; i += 2 {
				sum += i

				ch <- sum
			}

		}

	}()

	for n := range ch {
		fmt.Println(n)
	}

}

func ex5(p int) {
	ch := make(chan int)

	go func() {
		defer close(ch)

		if p%2 == 0 {
			for i := 0; i <= p; i += 2 {
				ch <- i
			}
		} else {
			p++ // Ensure p is even
			for i := 0; i <= p; i += 2 {
				ch <- i
			}
		}
	}()

	for n := range ch {
		fmt.Println(n)
	}
}

// has issues
func ex6(k string) {

	ch := make(chan []string)

	var l []string

	go func() {

		defer close(ch)
		for i := len(k) - 1; i < len(k); i-- {

			l = append(l, string(k[i]))
			result := make([]string, len(l))

			ch <- result

		}

	}()
	for n := range ch {
		fmt.Print(n)
	}

}

func ex7(k []int) <-chan []int {
	ch := make(chan []int)

	go func() {
		defer close(ch)
		var s []int
		for i := 0; i < len(k); i++ {
			s = append(s, 2*k[i])
			// Send a copy of the slice
			result := make([]int, len(s))
			copy(result, s)
			ch <- result
		}
	}()

	return ch
}

func main() {
	/*
			ex1()
			fmt.Println("first exe is done ")

			ex2()
			fmt.Println("second exe is done ")

		ex3(7)
		fmt.Println("three exe is done ")*/

	p := []int{4, 5, 3, 6, 3, 6, 3}
	ex7(p)
	fmt.Println("three exe is done ")

	slice := []int{9, 4, 7, 2, 1, 5, 8, 3, 6}

	resultCh := make(chan []int)

	go func() {
		sort.Ints(slice)
		resultCh <- slice
	}()

	sortedSlice := <-resultCh

	fmt.Println("Sorted Slice:", sortedSlice)
}
