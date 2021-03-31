package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand" // not for security critical thing, that's why to use math/rand instead of crypt/rand
	"os"
	"time"
)

var (
	words = []string{
		"abc",
		"bird",
		"lion",
		"tiger",
		"elephant",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func pickWord() string {
	return words[rand.Intn(len(words)-1)]
}

/*
標準出力に英単語を出す（出すものは自由）
標準入力から1行受け取る
制限時間内に何問解けたか表示する
*/

var (
	currentWord string
	successCont int = 0
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("success: %d\n", successCont)
			os.Exit(0)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	currentWord = pickWord()
	fmt.Println(currentWord)
	for scanner.Scan() {
		if scanner.Text() == currentWord {
			successCont++
		}

		currentWord = pickWord()
		fmt.Println(currentWord)
	}
}
