package main

import (
	"fmt"

	lib "github.com/barnakun/migratowl-bench-adv-inj"
)

// Hello greets name.
func Hello(name string) string { return lib.Greet(name) }

func main() { fmt.Println(Hello("world")) }
