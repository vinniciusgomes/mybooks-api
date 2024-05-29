package main

import "mybooks/internal/infrastructure/api"

func main() {
	if err := api.StartServer(); err != nil {
		panic(err)
	}
}
