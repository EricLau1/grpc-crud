package main

import "grpc-crud/server"

func main() {
	server.Run(":50000")
}
