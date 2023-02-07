package pkg1

import (
	"fmt"

	_ "github.com/bigwhite/prog-init-order/pkg2"
	_ "github.com/bigwhite/prog-init-order/pkg3"
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
		fmt.Println("pkg1: const c1 has been initialized")
	}
	if c2 != "" {
		fmt.Println("pkg1: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg1: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg1: first init func invoked")
}

func init() {
	fmt.Println("pkg1: second init func invoked")
}

func main() {
	// do nothing
}
