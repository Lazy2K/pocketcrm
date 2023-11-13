package main

import (
	"pocketcrm/packages/api"
	"pocketcrm/packages/database"
)

func initPocketcrm() {
	
}

func main() {
	// Check for first time start-up
	database.Connect()
	const PORT = ":8000" // Get this from a config file later or smthn for easy setup on user vms
	api.StartServer(PORT)
}
