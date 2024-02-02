package main

import (
	"pocketcrm/packages/api"
	"pocketcrm/packages/dashboard"
	"pocketcrm/packages/database"
	"sync"
)

func initPocketcrm() {
	
}

func main() {
	// Check for first time start-up
	database.Connect()
	const API_PORT = ":8000" // Get this from a config file later or smthn for easy setup on user vms
	const DASH_PORT = ":80" // Or even run them on the same port but have them use different paths

	group := new(sync.WaitGroup)
	group.Add(2)

	// Annon function to start API server
	go func() {
		api.StartServer(API_PORT)
		group.Done()
	}()

	// Annon function to start the Dashboard server
	go func() {
		dashboard.StartServer(DASH_PORT)
		group.Done()
	}()

	// Wait group will wait terminaiton of both servers
	group.Wait()
}
