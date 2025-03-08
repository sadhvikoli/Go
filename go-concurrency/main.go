package main

import (
	"fmt"
	"mymodule/practice"
	"sync"
	"time"
)

func printMessage() {
	fmt.Println("Hello from Goroutine!")
}

func printNumber(n int) {
	fmt.Println("Number:", n)
}

func printMessag(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that goroutine is done
	fmt.Println("Goroutine:", id)
}

func main() {
	go printMessage()       // Starts a new Goroutine
	time.Sleep(time.Second) // Give Goroutine time to execute
	fmt.Println("Main function finished!")
	// 	go functionName() → Runs a function concurrently.
	// Goroutines run independently from the main function.
	// Use time.Sleep (for now) to let the goroutine finish.

	// Running Multiple Goroutines
	for i := 1; i <= 5; i++ {
		go printNumber(i) // Launching multiple goroutines
	}
	time.Sleep(time.Second)
	// Output may vary due to concurrency

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add one goroutine to wait for
		go printMessag(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All Goroutines Finished!")
	// 	wg.Add(n) → Tells WaitGroup to wait for n goroutines.
	// wg.Done() → Signals that a goroutine has completed.
	// wg.Wait() → Blocks until all goroutines complete.

	ch := make(chan string) // Create a channel

	go func() {
		ch <- "Hello from Goroutine!" // Send data to channel
	}()

	msg := <-ch // Receive data from channel
	fmt.Println(msg)
	// 	ch := make(chan type) → Creates a channel.
	// ch <- value → Sends data into the channel.
	// <-ch → Receives data from the channel.

	done := make(chan bool)
	go worker(done)

	<-done // Wait for signal
	fmt.Println("Worker Done!")
	//Worker goroutine signals the main function using a channel.

	// ch := make(chan int) // Unbuffered channel

	// go func() {
	//     ch <- 100 // Blocks until received
	// }()

	// fmt.Println(<-ch) // Receives and prints 100

	// ch := make(chan int, 2) // Buffered with size 2

	// ch <- 10
	// ch <- 20

	// fmt.Println(<-ch) // 10
	// fmt.Println(<-ch) // 20
	// Buffered channels allow sending without immediate receiving, up to their capacity.

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
	// 	select waits for any channel to send data.
	// Whichever channel sends first, gets executed.

	practice.ExecuteFunc()
}

func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true // Signal completion
}
