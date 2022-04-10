package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func(d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}


type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.ListenAndServe("localhost:8080", db)
}

