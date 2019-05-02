package client

import "fmt"

type ClientRequestHandlerImpl struct{}

func (ClientRequestHandlerImpl) Send(msg []byte) {
	fmt.Println("TESTE TESTE")
}
