package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/tsconn23/event-checksum/cmd/mock"

	"github.com/tsconn23/event-checksum/events"
)

func main() {
	var useBinary bool
	var postEvent bool

	flag.BoolVar(&useBinary, "b", false, "Create event with binary payload")
	flag.BoolVar(&postEvent, "p", false, "Post resulting event to EdgeX")
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
	if !postEvent {
		checksum := sha256.Sum256(data)
		fmt.Printf("checksum: %x\r\n", checksum)
		fmt.Printf("data:\r\n%v\r\n", data)
	} else {
		params := types.EndpointParams{
			ServiceKey:  clients.CoreDataServiceKey,
			Path:        clients.ApiEventRoute,
			UseRegistry: false,
			Url:         "http://localhost:48080" + clients.ApiEventRoute,
			Interval:    30000,
		}


		client := coredata.NewEventClient(params, mock.MockEndpoint{})

		ctx := context.WithValue(context.Background(),clients.ContentType, clients.ContentTypeCBOR)
		id, err := client.AddBytes(data, ctx)
		if err != nil {
			fmt.Printf("error %s\r\n", err.Error())
			return
		}
		fmt.Printf("new id %s\r\n", id)
	}
}
