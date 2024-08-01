package main

import (
	"sync"
	"time"
)

func worker(input chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(3 * time.Second)
	msg := <-input
	println(msg)
}

func main() {

	input := make(chan string)

	var wg sync.WaitGroup
	wg.Add(10)
	go worker(input, &wg)

	input <- "hello"
	input <- "world"
	input <- "foo"

	wg.Wait()
}
