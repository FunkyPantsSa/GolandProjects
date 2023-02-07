package main

import (
	"fmt"

	_ "github.com/bigwhite/prog-init-order/pkg1"
	_ "github.com/bigwhite/prog-init-order/pkg2"
)

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized")
	}
	if c2 != "" {
		fmt.Println("main: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("main: first init func invoked")
}

func init() {
	fmt.Println("main: second init func invoked")
}

func main() {
	// do nothing
}
