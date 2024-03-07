package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers(ch chan int) {
	for i := 0; i < 3; i++ {
		randomNum := rand.Intn(50)
		ch <- randomNum
	}
	close(ch)
}

func main() {
	randomNumChannel := make(chan int)

	go generateRandomNumbers(randomNumChannel)

	for i := 0; i < 3; i++ {
		randomNum := <-randomNumChannel
		fmt.Printf("Received random number: %d\n", randomNum)
	}

	fmt.Println("All numbers received. Exiting...")
}
