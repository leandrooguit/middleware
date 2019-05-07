package dist

import (
	"fmt"
	"middleware/lib/infra/server"
)

type InvokerImpl struct {}

func (InvokerImpl) Invoke() (err error) {
	s := server.NewServerRequestHandlerImpl("1234")

	fmt.Println("Chegou")

	for {

		msgToBeUnmarshalled := s.Receive()

		msgReturned, err := Unmarshall(msgToBeUnmarshalled)

		if err != nil {
			return err
		}

		msgReturned.Body.ReplyHeader = ReplyHeader{}
		msgReturned.Body.ReplyBody = nil // todo implementar chamada do objeto remoto

		var bytes []byte
		bytes, err = Marshall(msgReturned)

		if err != nil {
			return err
		}

		s.Send(bytes)
	}

	return nil
}