package server

import "net"

type ServerRequestHandler interface {
	GetConn() net.Conn
	Receive(conexao net.Conn) []byte
	Send(msg []byte)
}

