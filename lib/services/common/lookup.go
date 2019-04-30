package common

type clientProxy struct {
	ip       string
	port     int
	protocol string // Quando o Client Request Handler usa mais de um protocolo, a AOR também inclui a identificação do protocolo.
	objectId int
}

type namingRecord struct {
	serviceName string
	clientProxy clientProxy
}

type lookup interface {
	bind(nr namingRecord)
	lookup(serviceName string) clientProxy
}
