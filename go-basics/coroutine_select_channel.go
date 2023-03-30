package main

import (
	"fmt"
	"time"
)

// Synchronous/blocking communication between forked and main goroutine
func goroutine() {
	channel := make(chan string)

	go func() {
		channel <- "data"
	}()

	msg := <-channel
	fmt.Println(msg)
}

// Data comes from any one of the channels will be printed
func goroutineWithSelect() {
	channel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		channel <- "data"
	}()

	go func() {
		anotherChannel <- "another data"
	}()

	select {
	case msg := <-channel:
		fmt.Println(msg)
	case msg := <-anotherChannel:
		fmt.Println(msg)
	}
}

// Bufferred channel
func bufferredChannel() {
	channel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case channel <- char:
		}
	}

	close(channel)

	for char := range channel {
		fmt.Println(char)
	}
}

// Auto curoff after N seconds
func cutoffAfterNSecs() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Running...")
			}
		}
	}()

	time.Sleep(time.Second * 4)
}

// Auto curoff channel only when close(channel) is invoked
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Running...")
		}
	}
}

func cutoffWhenDoneChannel() {
	done := make(chan bool)
	doWork(done)
	time.Sleep(time.Second * 1)
	close(done)
}
