package client

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"
)

type ClientRequestHandlerImpl struct {
	Host string
	Port int
	conect net.Conn
}

func NewClientRequestHandlerImpl(host string, port int) *ClientRequestHandlerImpl {

	address := host + ":" + strconv.Itoa(port)

	conexao, erro1 := net.Dial("tcp", address)

	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	return &ClientRequestHandlerImpl{Host: host, Port: port, conect: conexao}
}

func (crh ClientRequestHandlerImpl) Receive() (msg []byte) {
	return nil
	//panic("implement me")
}

func (crh ClientRequestHandlerImpl) Send(msg []byte) {

	_, erro2 := crh.conect.Write(msg)

	fmt.Println("Mensagem enviada de forma ficticia")
}
