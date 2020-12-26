package main

import (
	"bufio"
	"log"
	"os"
	"os/user"
)

func writeHostToConfig(h *host) {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Couldn't get current user: %s", err.Error())
		return
	}
	file, _ := os.OpenFile(u.HomeDir + "/.ssh/config", os.O_APPEND|os.O_WRONLY, 0774)
	writer := bufio.NewWriter(file)
	linesToWrite := h.getWritableHost()
	for _, line := range linesToWrite {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Couldn't write to file: %s", err.Error())
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatalf("Couldn't flush writer: %s", err.Error())
	}
}