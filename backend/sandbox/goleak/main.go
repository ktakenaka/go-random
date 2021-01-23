package main

import (
	"log"
	"time"
)

func leak() {
	ch := make(chan int)
	go func() {
		val := <-ch
		log.Println(val)
	}()
}

func main() {
	doneCh := make(chan struct{})
	for i := 1; i <= 3; i++ {
		i := i
		go forSelectPattern(i, doneCh)
	}
	time.Sleep(10 * time.Second)
	close(doneCh)
}

// 終了のシグナルが来るまで、繰り返し作業を行うパターンで便利
func forSelectPattern(n int, doneCh <-chan struct{}) {
	for {
		select {
		case <-time.Tick(time.Duration(n) * time.Second):
			log.Printf("hello, %d desu. waiting", n)
		case <-doneCh:
			log.Printf("finished %d", n)
			return
		}
	}
}
