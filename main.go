package main

import "wux4an/gats/services"

func main() {
	server := services.Server()
	server.Run()
}
