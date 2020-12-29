package main

import (
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
			break
		}
	} else {
		// Get a host from the user.
		h := getHostFromUser()
		// Get all hosts
		hostList := getAllHosts()
		// Append or replace the host to the list.
		hostList[h.name] = h
		// Write the host to the config file.
		writeHosts(hostList)
	}
}

// Removes a single host based on the arguments provided.
func removeHost(args []string) {
	var hostToDelete string
	if len(args) >= 3 {
		hostToDelete = args[2]
	} else {
		hostToDelete = getHostNameFromUser()
	}
	hosts := getAllHosts()
	delete(hosts, hostToDelete)
	writeHosts(hosts)
}

// Writes a map of hostst to the config file. The current hosts will be removed.
func writeHosts(hosts map[string]*host) {
	// Remove all current hosts from the file.
	clearHostsFromConfig()
	// Add ever single host to the config file.
	for _, h := range hosts {
		writeHostToConfig(h)
	}
}
