package main

import (
	"fmt"
	"example.com/tempconv"
)

func main() {
	var flag tempconv.CelsiusFlag
	flag.Set("212F")
	fmt.Println(flag)
}
