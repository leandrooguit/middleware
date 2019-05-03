package dist

type Invocation interface {
	ObjectId() int
	SetObjectId(objectId int)
	IpAddress() string
	SetIpAddress(ipAddress string)
	PortNumber() int
	SetPortNumber(portNumber int)
	OperationName() string
	SetOperationName(operationName string)
	Parameters() [2]string
	SetParameters(parameters [2]string)
}

type Termination interface {
	Result() interface{}
}

type Requestor interface {
	Invoke(inv Invocation) Termination
}
