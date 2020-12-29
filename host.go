package main

import (
	"fmt"
	"log"
	"regexp"
)

type host struct {
	name              string
	hostName          string
	user              string
	additionalOptions map[string]string
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

func getAllHosts() map[string]*host {
	fileContent, err := getFileContent()
	if err != nil {
		log.Fatalf("Error geting the file content: %s", err.Error())
	}
	var hostList map[string]*host
	var h *host
	nameRegex := regexp.MustCompile("(?<=Host\s)\w+")
	optionRegex := regexp.MustCompile("(?<=\s)\w+")
	for _, lineValue := range fileContent {
		if result, _ := regexp.MatchString("(?<=Host\s)\w+"); result {
			if h != nil {
				hostList[h.name] = h
			}
			h = host{}
			h.name = nameRegex.FindString(lineValue)
		}
		else if result, _ := regexp.MatchString("^\s+\w+\s+\w+\s*$"); result {
			results := optionRegex.FindAllString()
			if len(results) == 2 {
				optionName := results[0]
				optionValue := results[1]
				h.additionalOptions[optionName] = optionValue
			}
		}
	}
	return hostList
}
