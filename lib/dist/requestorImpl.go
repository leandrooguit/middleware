package dist

import (
	"middleware/lib/infra/client"
	"unsafe"
)

// Implements Invocation
type InvocationImpl struct {
	objectId      int
	ipAddress     string
	portNumber    int
	operationName string
	parameters    [2]string
}

func (i InvocationImpl) ObjectId() int {
	return i.objectId
}

func (i InvocationImpl) SetObjectId(objectId int) {
	i.objectId = objectId
}

func (i InvocationImpl) IpAddress() string {
	return i.ipAddress
}

func (i InvocationImpl) SetIpAddress(ipAddress string) {
	i.ipAddress = ipAddress
}

func (i InvocationImpl) PortNumber() int {
	return i.portNumber
}

func (i InvocationImpl) SetPortNumber(portNumber int) {
	i.portNumber = portNumber
}

func (i InvocationImpl) OperationName() string {
	return i.operationName
}

func (i InvocationImpl) SetOperationName(operationName string) {
	i.operationName = operationName
}

func (i InvocationImpl) Parameters() [2]string {
	return i.parameters
}

func (i InvocationImpl) SetParameters(parameters [2]string) {
	i.parameters = parameters
}

// Implements Termination
type TerminationImpl struct {
	result interface{}
}

func (t TerminationImpl) Result() interface{} {
	panic("implement me")
}

// Implements requestor
type RequestorImpl struct{}

func (RequestorImpl) Invoke(inv Invocation) (t Termination, err error) {

	crh := client.NewClientRequestHandlerImpl(inv.IpAddress(), inv.PortNumber())

	msg := MessageImpl{HeaderImpl{"MID", "0.1", true, 0, 0}, BodyImpl{string()}} // inv.IpAddress() / strconv.Itoa(inv.PortNumber()) / inv.OperationName())

	var bytes []byte
	bytes, err = Marshall(msg)
	if err != nil {
		return nil, err
	}

	crh.Send(bytes)

	msgReturned, err := Unmarshall(crh.Receive())
	if err != nil {
		return nil, err
	}

	t = TerminationImpl{msgReturned.Body().ReplyBody()}

	return t, err
}
