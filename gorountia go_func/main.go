package main

import (
	"fmt"
)

/*
func sum(wg *sync.WaitGroup) {
	defer wg.Done() // Mark the current goroutine as done when finished
	k := 100
	for i := 0; i < k; i++ {
		fmt.Println(i)
	}

}

func main() {
	var wg sync.WaitGroup

	// Add one goroutine to the wait group
	wg.Add(1)

	// Start the sum function as a goroutine and pass the wait group
	go sum(&wg)

	// Wait for all goroutines in the wait group to finish
	wg.Wait()

	fmt.Println("Sum completed")
}
*/

func fibonacci(n int) []int {
	var fib []int

	if n <= 0 {
		return fib
	}

	fib = append(fib, 0)
	if n == 1 {
		return fib
	}

	fib = append(fib, 1)
	if n == 2 {
		return fib
	}
	package main
	
	import (
		"fmt"
		"net/http"
		"os"
		"strings"
	
		"golang.org/x/net/html"
	)
	
	func main() {
	
		getinfo("https://en.wikipedia.org/wiki/Iron_Man")
	
	}
	
	func getinfo(hhtplink string) {
		response, err := http.Get(hhtplink)
		if err != nil {
			fmt.Println("error getting information:", err)
			return
		}
		defer response.Body.Close()
	
		filename := "iron.txt2"
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
	
		// Parse the HTML content
		z := html.NewTokenizer(response.Body)
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
				// End of the document, we're done
				fmt.Println("File created:", filename)
				return
			case tt == html.TextToken:
				// Text token found, extract and write to the file
				text := strings.TrimSpace(string(z.Text()))
	
				if len(text) > 0 {
					_, err := file.WriteString(text + "\n")
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}

func main() {
	n := 10 // You can change the value of n to calculate the Fibonacci sequence up to n terms
	fib := fibonacci(n)
	fmt.Println("Fibonacci Sequence:", fib)
}
