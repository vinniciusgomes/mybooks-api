package main

import "mybooks/internal/infrastructure/api"

// main is the entry point of the Go program.
//
// It calls the StartServer function from the api package to start the server.
// If an error occurs during the server startup, it panics.
func main() {
	if err := api.StartServer(); err != nil {
		panic(err)
	}
}
