package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

func getNumbers(in1 <-chan int, in2 <-chan int) (int, int) {
	var wg sync.WaitGroup
	wg.Add(2)
	var x1, x2 int
	getX1 := func() {
		x1 = <-in1
		defer wg.Done()
	}
	getX2 := func() {
		x2 = <-in2
		defer wg.Done()
	}

	go getX1()
	go getX2()
	wg.Wait()

	fmt.Printf("x1 = %v, x2 = %v\n", x1, x2)
	return x1, x2
}

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		mutex.Lock()
		for i := 0; i < n; i++ {
			x1, x2 := getNumbers(in1, in2)
			result := f(x1) + f(x2)
			fmt.Printf("result = %v, x1 = %v, x2 = %v\n", result, x2, x1)
			out <- result
		}
		mutex.Unlock()
	}()
}

func plusOne(num int) int {
	time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
	return num + 1
}

func main() {
	var n = 12
	channel1 := make(chan int)
	channel2 := make(chan int)
	channelResult := make(chan int)

	Merge2Channels(plusOne, channel1, channel2, channelResult, n)

	for i := 0; i < 3; i++ {
		channel1 <- i + 1
		channel2 <- i + 2
		fmt.Println(<-channelResult)
	}
	for i := 0; i < 3; i++ {
		channel2 <- i + 2
		channel1 <- i + 1
		fmt.Println(<-channelResult)
	}

	for i := 0; i < 3; i++ {
		channel2 <- i + 2
		channel1 <- i + 1
	}

	for i := 0; i < 3; i++ {
		channel1 <- i + 1
	}
	for i := 0; i < 3; i++ {
		channel2 <- i + 2
		fmt.Println(<-channelResult)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-channelResult)
	}
}
