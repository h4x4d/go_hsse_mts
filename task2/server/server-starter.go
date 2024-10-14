package main

import (
	"task2/server/args"
	"task2/server/server"
)

func main() {
	port := args.GetPort()

	serv := server.NewServer(port)
	serv.Start()
}
