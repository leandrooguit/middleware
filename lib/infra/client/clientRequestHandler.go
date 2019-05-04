package client

type ClientRequestHandler interface {
	Send(msg []byte)
	Receive() (msg []byte)
}
