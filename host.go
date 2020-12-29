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

func getAllHosts() map[int]*host {
	fileContent, err := getFileContent()
	if err != nil {
		log.Fatalf("Error geting the file content: %s", err.Error())
		return nil
	}
	var hostList = make(map[int]*host)
	var h *host
	nameRegex := regexp.MustCompile("Host\\s+\\S+")
	nameRegexReplace := regexp.MustCompile("Host\\s+")
	optionRegex := regexp.MustCompile("\\s+\\S+")
	optionRegexReplace := regexp.MustCompile("\\s+")
	for i := 0; i < len(fileContent); i++ {
		lineValue := fileContent[i]
		if result, _ := regexp.MatchString("Host\\s+\\S+", lineValue); result {
			if h != nil {
				hostList[len(hostList)] = h
			}
			h = &host{}
			h.additionalOptions = make(map[string]string)
			name := nameRegex.FindString(lineValue)
			h.name = nameRegexReplace.ReplaceAllString(name, "")
		} else if result, _ := regexp.MatchString("^\\s+\\w+\\s+\\S+$", lineValue); result {
			results := optionRegex.FindAllString(lineValue, -1)
			if len(results) == 2 {
				optionName := optionRegexReplace.ReplaceAllString(results[0], "")
				optionValue := optionRegexReplace.ReplaceAllString(results[1], "")
				h.additionalOptions[optionName] = optionValue
			}
		}
	}
	if h != nil {
		hostList[len(hostList)] = h
	}
	return hostList
}
