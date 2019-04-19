package main

import (
	"crypto/sha256"
	"flag"
	"fmt"

	"github.com/tsconn23/event-checksum/events"
)

func main() {
	var useBinary bool

	flag.BoolVar(&useBinary, "b", false, "Create event with binary payload")
	flag.Parse()

	var data []byte
	var err error
	if useBinary {
		data, err = events.NewBinaryEvent()
	} else {
		data, err = events.NewBasicEvent()
	}

	if err != nil {
		fmt.Sprintln(err.Error())
		return
	}
	checksum := sha256.Sum256(data)
	fmt.Printf("checksum: %x\r\n", checksum)
}
