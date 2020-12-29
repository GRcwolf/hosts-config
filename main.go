package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argumnets := os.Args
	if len(argumnets) > 1 && argumnets[1] != "create" {
		switch argumnets[1] {
		case "remove":
			removeHost(argumnets)
			break
		default:
			log.Fatal("Invalid argument")
		}
	} else {
		// Get a host from the user.
		h := getHostFromUser()
		// Write the host to the config file.
		writeHostToConfig(h)
	}
}

func removeHost(args []string) {
	var hostToDelete string
	if len(args) >= 3 {
		hostToDelete = args[2]
	} else {
		fmt.Println("Which host should be deleted?")
		fmt.Scanln(&hostToDelete)
	}
	hosts := getAllHosts()
	delete(hosts, hostToDelete)
	writeHosts(hosts)
}

func writeHosts(hosts map[string]*host) {
	clearHostsFromConfig()
	for _, h := range hosts {
		writeHostToConfig(h)
	}
}
