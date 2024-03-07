package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(ch chan<- int) {
	randomNum := rand.Intn(100) + 1
	ch <- randomNum
}

func main() {
	randomNumChannel := make(chan int)

	go generateRandomNumber(randomNumChannel)

	for {
		select {

		case randomNumber := <-randomNumChannel:

			time.After(3 * time.Second)
			fmt.Printf("Random number received: %d\n", randomNumber)

			return
		}

	}
}
