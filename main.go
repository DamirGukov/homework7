package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	c2 := make(chan int)
	sum := 0

	rand.Seed(time.Now().UnixNano())

	go func() {
		for i := 0; i < 3; i++ {
			n := rand.Intn(100)
			fmt.Println("Random number:", n)
			c <- n
		}
		close(c)
	}()

	go func() {
		for value := range c {
			sum += value
		}
		sum = sum / 3
		c2 <- sum
	}()

	go func() {
		answer := <-c2
		fmt.Println("Average value:", answer)
	}()
	time.Sleep(time.Second)

}
