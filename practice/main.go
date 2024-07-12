package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	var mychan = make(chan int, 1)
	wg.Add(1)
	go Value(mychan, &wg)
	mychan <- 10
	defer wg.Wait()
}

func Value(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Hello", <-ch)
	ch <- 20
	wg.Add(1)
	go Hello(ch, wg)
	defer wg.Done()
}

func Hello(ch chan int, wg *sync.WaitGroup) {
	_, err := fmt.Println("Hi", <-ch)
	if err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
}
