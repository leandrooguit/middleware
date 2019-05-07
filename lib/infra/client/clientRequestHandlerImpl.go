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
}

func (crh ClientRequestHandlerImpl) Receive() (msg []byte) {
	return nil
	//panic("implement me")
}

func (crh ClientRequestHandlerImpl) Send(msg []byte) {
	address := crh.Host + ":" +strconv.Itoa(crh.Port)

	conexao, erro1 := net.Dial("tcp", address)

	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	_, erro2 := conexao.Write(msg)

	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	fmt.Println("Mensagem enviada de forma ficticia")
}
