package client

import (
	"fmt"
)

type ClientRequestHandlerImpl struct{}

func (crh ClientRequestHandlerImpl) Receive() (msg []byte) {
	return nil
	//panic("implement me")
}

func (crh ClientRequestHandlerImpl) Send(msg []byte) {
	fmt.Println("Mensagem enviada de forma ficticia")
}
