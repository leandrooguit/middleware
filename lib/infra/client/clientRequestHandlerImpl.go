package client

import (
	"fmt"
	"net"
	"os"
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
	_, err := crh.conect.Read(msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	return msg
}

func (crh ClientRequestHandlerImpl) Send(msg []byte) {
	_, erro2 := crh.conect.Write(msg)
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}
}
