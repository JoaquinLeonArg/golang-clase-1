package main

import "github.com/joaquinleonarg/go-pokemon/src/api"

func main() {
	// TODO: Loaad env from file to decide kind of service
	api.StartServer(8082, "TODO")
}
