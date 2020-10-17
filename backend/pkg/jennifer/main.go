package main

import (
	"io/ioutil"

	//nolint:golint,stylecheck
	. "github.com/dave/jennifer/jen"
)

// go run pkg/jennifer/main
func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)

	if err := ioutil.WriteFile("trial.go", []byte(f.GoString()), 0600); err != nil {
		panic(err)
	}
}
