package main

import (
	"fmt"
	"time"
)

/*
func sendRandomNumber() (chan int, <-chan string) {

	randomCh := make(chan int)
	readableCh := make(chan string)

	go func() {

		defer close(randomCh)
		defer close(readableCh)

		for {

			select {
			case randomCh <- rand.Intn(100):
				time.Sleep(time.Second * 10)
				readableCh <- "Number sent "
			case <-time.After(time.Second): //Timeout after 1 second

				readableCh <- "Timeout occurred"

				return

			}

		}

	}()

	return randomCh, readableCh

}

func main() {

	randomch, readlech := sendRandomNumber()

	for {

		select {

		case num := <-randomch:
			time.Sleep(time.Second * 4)
			fmt.Println("received random number:", num)
		case status := <-readlech:

			fmt.Println("status:", status)

		}

	}

}



+===================================================================
==================================================
===========================+++++++++++++++++++++++++++++++++++=


/*
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchURL(k string) {

	urlrespon := make(chan string)

	resp, err := http.Get(k)
	if err != nil {
		fmt.Println("error fetching %s: %s\n", k, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	body = []byte(<-urlrespon)

}



func main(){

now:=time.Now()

urls:=[]string{

"https://www.google.com",
"https://www.hithub.com",

go _,url:=range urls{

	go fetchURL(url )
}



}

fmt.Println("done in:",time.Since(now).Second)


}
===============================================================================
==================================================================================

==================================================================================
==================================================================
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

funcurlResponses fetchURL(url string, status chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading body for %s: %s\n", url, err)
		return
	}

	ch <- string(body)
	s<-string(err)
}

func main() {
	now := time.Now()

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
	}

	urlResponses := make(chan string)


	for _, url := range urls {
		go fetchURL(url, urlResponses)
	}

	for  {
		select{

		case url:=<-urlResponses:





		}



	}




	fmt.Println("done in:", time.Since(now).Seconds())
}

/*
type status struct{}

func worker(id int, done chan struct{}) {
	defer func() {
		done <- struct{}{} // signal that goroutine is done
	}()

	fmt.Printf("worker %d started\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	numWorkers := 3
	done := make(chan struct{}, numWorkers) // buffered channel to prevent goroutine leak

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, done)
	}

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		select {
		case <-done:
			fmt.Println("received signal from a worker. A worker is done")
		case <-time.After(2 * time.Second):
			fmt.Println("timeout reached. Exiting...")
			return
		}
	}
}
*/