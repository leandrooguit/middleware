package main

import "middleware/lib/dist"

func main() {
	r := new(dist.RequestorImpl)

	var parameters [2]string
	parameters[0] = "Pedra"
	parameters[1] = "Tesoura"

	i := dist.InvocationImpl{ObjectId: 1000, IpAddress: "127.0.0.1:1234", PortNumber: 1234, OperationName: "jankenpo", Parameters: parameters}

	r.Invoke(i)
}
