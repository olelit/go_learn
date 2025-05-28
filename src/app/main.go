package main

import (
	"learn/go/internal/database"
	"learn/go/internal/routing"
	"net/http"
)

func main() {
	database.Connect()
	r := routing.Route()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
