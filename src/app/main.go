package main

import (
	"learn/go/internal/routing"
	"net/http"
)

func main() {
	var r = routing.Route()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
