package main

import (
	"github.com/wesley-gomes-sje/api-go-rest/database"
	"github.com/wesley-gomes-sje/api-go-rest/routes"
)

func main() {

	database.ConnectionDb()
	routes.HandleRequests()
}
