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
	Parameters() []interface{}
	SetParameters(parameters []interface{})
}

type Termination interface {
	Result() interface{}
}

type Requestor interface {
	Invoke(inv Invocation) (t Termination, err error)
}
