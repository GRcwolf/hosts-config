package main

import (
	"bufio"
	"log"
	"os"
	"os/user"
)

// Writes thr host to the config file.
func writeHostToConfig(h *host) {
	// Load the current user.
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Couldn't get current user: %s", err.Error())
		return
	}
	// Open file in order to later append the host.
	file, err := os.OpenFile(u.HomeDir+"/.ssh/config", os.O_APPEND|os.O_WRONLY, 0774)
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
		return
	}
	// Make sure the file will be closed.
	defer closeFile(file)
	writer := bufio.NewWriter(file)
	// Get the lines to write.
	linesToWrite := h.getWritableHost()
	for _, line := range linesToWrite {
		// Write each line.
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

// Closes the os.File that is passes as parameter and shows an error if this isn't possible.
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalf("Couldn't close file: %s", err.Error())
	}
}

func getFileContent() (map[int]string, error) {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Couldn't get current user: %s", err.Error())
		return nil, err
	}
	file, err := os.OpenFile(u.HomeDir+"/.ssh/config", os.O_RDONLY, 0774)
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
		return nil, err
	}
	defer closeFile(file)
	reader := bufio.NewScanner(file)
	var lines map[int]string
	for reader.Scan() {
		lines[len(lines)] = reader.Text()
	}
	return lines, nil
}
