package main

import "os"

func main() {
	argumnets := os.Args
	if len(argumnets) > 1 && argumnets[1] != "create" {

	} else {
		// Get a host from the user.
		h := getHostFromUser()
		// Write the host to the config file.
		writeHostToConfig(h)
	}
}
