package mock

import "github.com/edgexfoundry/go-mod-core-contracts/clients/types"

type MockEndpoint struct {

}

func (e MockEndpoint) Monitor(params types.EndpointParams, ch chan string) {
	//do nothing
}