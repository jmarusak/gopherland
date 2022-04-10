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
	switch req.URL.Path {
	case "/list":
		for item, price := range d {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")

		price, ok := d[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "not found item: %s\n", item)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.ListenAndServe("localhost:8080", db)
}

