package main

import (
	"pocketcrm/packages/api"
	"pocketcrm/packages/database"
	"pocketcrm/packages/web"
	"sync"
)

func initPocketcrm() {
	
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	const API_PORT = ":8000"
	const WEB_PORT = ":8001" // Http -- upgrade to https or smthn with SSL somehow
	// Check for first time start-up
	database.Connect()
	 // Get this from a config file later or smthn for easy setup on user vms


	 go func() {
		api.StartServer(API_PORT)
		wg.Done()
	 }()

	 go func() {
		web.StartServer(WEB_PORT)
		wg.Done()
	 }()

	 wg.Wait()

	
}
