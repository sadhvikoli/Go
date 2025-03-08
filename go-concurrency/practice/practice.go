package practice

import (
	"fmt"
	"sync"
	"time"
)

// Create a goroutine that prints numbers from 1 to 10 concurrently.
func printNumber(n int) {
	fmt.Println("Number:", n)
}

func printMessag(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that goroutine is done
	fmt.Println("Goroutine:", id)
}

func ExecuteFunc() {
	for i := 1; i <= 10; i++ {
		go printNumber(i) // Launching multiple goroutines
	}
	time.Sleep(time.Second)

	// Modify the above program to use WaitGroup instead of time.Sleep().
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add one goroutine to wait for
		go printMessag(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Goroutines Finished!")

	// Implement a channel to pass a string message from one goroutine to another.
	messageChannel := make(chan string)

	go func() {
		messageChannel <- "Hello from Goroutine!"
	}()

	msg := <-messageChannel
	fmt.Println("Received:", msg)

	// Task 2: Buffered channel of size 3
	bufferedChan := make(chan string, 3)

	bufferedChan <- "First"
	bufferedChan <- "Second"
	bufferedChan <- "Third"

	fmt.Println("Buffered Channel Receives:")
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	// Select statement to listen to two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from channel 2"
	}()

	// Use select to receive from whichever channel responds first
	select {
	case msg1 := <-channel1:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}
}
