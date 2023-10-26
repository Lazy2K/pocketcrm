package main

import (
	"pocketcrm/api"
)

func main() {
	const PORT = ":8000" // Get this from a config file later or smthn for easy setup on user vms
	api.StartServer(PORT)
}
