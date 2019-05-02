package dist

type Invocation interface {
}

type Termination interface {
}

type Requestor interface {
	Invoke(inv Invocation) Termination
}
