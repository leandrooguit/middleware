package server

import (
	"fmt"
	"net"
	"os"
)

type ServerRequestHandlerImpl struct {
	Port string
	conect net.Conn
}

func NewServerRequestHandlerImpl(port string) *ServerRequestHandlerImpl {

	fmt.Println("Servidor aguardando conexões...")

	// Ouvindo na porta 1234 via protocolo tcp/ip
	ln, erro1 := net.Listen("tcp", ":"+ port)

	if erro1 != nil {
		fmt.Println(erro1)

		/* Neste nosso exemplo vamos convencionar que a saída 3 está reservada para erros de conexão.
		IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()

	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexão aceita...")

	return &ServerRequestHandlerImpl{Port: port, conect: conexao}
}

func (s ServerRequestHandlerImpl) Receive() []byte {

	msgFromClient := make([]byte, 1024)

	// recebe e desserializa request
	n, erro3 := s.conect.Read(msgFromClient)
	if erro3 != nil {
		fmt.Println(erro3)
		os.Exit(3)
	}

	return msgFromClient[:n]
}

func (s ServerRequestHandlerImpl) Send(msg []byte) {
	// Enviar mensagem processada de volta para o cliente
	_, error := s.conect.Write(msg)
	if error != nil{
		fmt.Println(error)
		os.Exit(3)
	}
}