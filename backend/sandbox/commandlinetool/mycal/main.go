package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	isLineNum := flag.Bool("n", false, "show line numbers")
	flag.Parse()
	fileNames := flag.Args()
	if len(fileNames) == 0 {
		log.Fatal("no fileNames")
	}

	for _, name := range fileNames {
		if *isLineNum {
			readFileWithLineNum(name)
		} else {
			readFile(name)
		}
	}
}

func readFile(name string) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFileWithLineNum(name string) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
