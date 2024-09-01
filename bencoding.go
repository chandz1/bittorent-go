package main

import (
	"fmt"
	"log"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Open the torrent file
	file, err := os.Open("test.torrent")
	logError(err)

	// Decode the torrent file
	data, err := bencode.Decode(file)
	logError(err)

	fmt.Printf("File: %v\n", data)
}
