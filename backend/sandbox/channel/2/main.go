package main

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

// errgroup.Groupはsync.WaitGroupと同じだが、並行実行のgoroutineとからエラーを受け取れる
// 内部的には、sync.Onceを使って、最初に受け取ったエラーのみを保持するようになっている
func main() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		n := i
		eg.Go(func() error {
			return do(n)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Print(err)
	}
}
func do(n int) error {
	if n%2 == 0 {
		return errors.New("err")
	}
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)
	return nil
}
