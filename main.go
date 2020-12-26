package main

func main() {
	// Get a host from the user.
	h := getHostFromUser()
	// Write the host to the config file.
	writeHostToConfig(h)
}