package main

import (
	"fmt"
	"os"
)

func main() {
	hosts := getAllHosts()
	argumnets := os.Args
	if len(argumnets) > 1 && argumnets[1] != "create" {
		hosts = getAllHosts()
		for i := 0; i < len(hosts); i++ {
			h := hosts[i]
			fmt.Println(h.name)
			writeHostToConfig(h)
		}
	} else {
		// Get a host from the user.
		h := getHostFromUser()
		// Write the host to the config file.
		writeHostToConfig(h)
	}
}
