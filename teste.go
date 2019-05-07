package main

import (
	"middleware/lib/dist"
)

func main() {
	teste := dist.InvokerImpl{}
	teste.Invoke()
}
