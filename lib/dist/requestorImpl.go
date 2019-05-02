package dist

import (
	"middleware/lib/infra/client"
	"strconv"
)

// Implements invocation
type InvocationImpl struct {
	ObjectId      int
	IpAddress     string
	PortNumber    int
	OperationName string
	Parameters    [2]string
}

// Implements Termination
type TerminationImpl struct{}

// Implements requestor
type RequestorImpl struct{}

func (RequestorImpl) Invoke(inv InvocationImpl) Termination {
	clientRequestHandler := new(client.ClientRequestHandlerImpl)
	msg := []byte(inv.IpAddress + "#" + strconv.Itoa(inv.PortNumber) + "#" + inv.OperationName)

	clientRequestHandler.Send(msg)
	return TerminationImpl{}
}
