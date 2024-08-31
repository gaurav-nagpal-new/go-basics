package main

import (
	"basics/test"
	"fmt"
	"os"
)

func main() {
	fmt.Println(test.Say(os.Args[1:]))
}
