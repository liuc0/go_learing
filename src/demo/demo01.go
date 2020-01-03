package main

import (
	"flag"
	"fmt"
)

var name  string
func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main()  {
	flag.Parse()
	hello(name)
}