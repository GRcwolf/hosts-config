package main

import (
	"fmt"
)

func getHostNameFromUser() string {
	hostName := ""
	var err error
	var code int
	for err != nil ||  hostName == ""  {
		fmt.Println("Enter the HostName:")
		code, err = fmt.Scanln(&hostName)
		if err != nil  {
			fmt.Printf("Got code %d, error %s", code, err)
		}
	}
	return hostName
}

func getUserFromUser() string {
	user := ""
	var err error
	var code int
	for err != nil ||  user == ""  {
		fmt.Println("Enter the user:")
		code, err = fmt.Scanln(&user)
		if err != nil  {
			fmt.Printf("Got code %d, error %s", code, err)
		}
	}
	return user
}

func getConfigName() string {
	configName := ""
	var err error
	var code int
	for err != nil ||  configName == ""  {
		fmt.Println("Enter the config name:")
		code, err = fmt.Scanln(&configName)
		if err != nil  {
			fmt.Printf("Got code %d, error %s", code, err)
		}
	}
	return configName
}

func getHostFromUser() *host {
	configName := getConfigName()
	hostName := getHostNameFromUser()
	user := getUserFromUser()
	return newHost(configName, hostName, user)
}