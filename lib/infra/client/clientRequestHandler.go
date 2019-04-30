package client

type clientRequestHandler interface {
	send(msg []byte)
	receive() (msg []byte)
}
