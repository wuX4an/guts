package main

import (
	"wux4an/gats/services"
	"wux4an/gats/services/bridges/database"
)

func main() {
	database.Create()
	server := services.Server()
	server.Run()
}
