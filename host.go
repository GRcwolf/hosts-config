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
func (h *host) getWritableHost() map[int]string {
	var writableLines = make(map[int]string)
	writableLines[len(writableLines)] = fmt.Sprintf("Host %s", h.name)
	if h.hostName != "" {
		writableLines[len(writableLines)] = fmt.Sprintf("\tHostName %s", h.hostName)
	}
	if h.user != "" {
		writableLines[len(writableLines)] = fmt.Sprintf("\tUser %s", h.user)
	}
	for name, value := range h.additionalOptions {
		writableLines[len(writableLines)] = fmt.Sprintf("\t%s %s", name, value)
	}
	writableLines[len(writableLines)] = ""
	return writableLines
}

// Gets all hosts from the cofig file.
func getAllHosts() map[string]*host {
	// Get the lines of the file.
	fileContent, err := getFileContent()
	if err != nil {
		log.Fatalf("Error geting the file content: %s", err.Error())
		return nil
	}
	var hostList = make(map[string]*host)
	var h *host
	// Regex to check for host line.
	nameRegex := regexp.MustCompile("Host\\s+\\S+")
	// Regex to replace the host line.
	nameRegexReplace := regexp.MustCompile("Host\\s+")
	// Regex for any host option.
	optionRegex := regexp.MustCompile("\\s+\\S+")
	// Replace regex for the host options.
	optionRegexReplace := regexp.MustCompile("\\s+")
	// Iterate through the lines in the correct order.
	for i := 0; i < len(fileContent); i++ {
		lineValue := fileContent[i]
		// Check if the line contains a host.
		if result, _ := regexp.MatchString("Host\\s+\\S+", lineValue); result {
			if h != nil {
				// Add the current host to the list.
				hostList[h.name] = h
			}
			// Create a new host.
			h = &host{}
			// Initiate the additional parameters.
			h.additionalOptions = make(map[string]string)
			name := nameRegex.FindString(lineValue)
			// Clean up the name.
			h.name = nameRegexReplace.ReplaceAllString(name, "")
		} else if result, _ := regexp.MatchString("^\\s+\\w+\\s+\\S+$", lineValue); result {
			results := optionRegex.FindAllString(lineValue, -1)
			// Make sure the option name and a value have been set.
			// This will cause problems if the value would contain whitespaces.
			if len(results) == 2 {
				optionName := optionRegexReplace.ReplaceAllString(results[0], "")
				optionValue := optionRegexReplace.ReplaceAllString(results[1], "")
				h.additionalOptions[optionName] = optionValue
			}
		}
	}
	// Add the last host to the config.
	if h != nil {
		hostList[h.name] = h
	}
	return hostList
}
