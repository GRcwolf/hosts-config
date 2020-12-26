package main

import "fmt"

type host struct {
	name string
	hostName string
	user string
}

// Creates a new host struct.
func newHost(name string, hostName string, user string) *host {
	return &host{name: name, hostName: hostName, user: user}
}

// Gets an array of string containing the single config lines.
func (h *host) getWritableHost() []string {
	writableLines := []string{
		fmt.Sprintf("Host %s", h.name),
		fmt.Sprintf("\tHostName %s", h.hostName),
		fmt.Sprintf("\tUser %s", h.user),
		"",
	}
	return writableLines
}