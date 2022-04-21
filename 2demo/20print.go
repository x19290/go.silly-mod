package main

import (
	"fmt"
	a "github.com/x19290/go.silly-mod"
	b "github.com/x19290/go.silly-mod/1testdata"
)

func Print() {
	fmt.Println(b.Version, a.Public())
}

func main() {
	Print()
}
