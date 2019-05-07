package dist

import (
	"fmt"
	"middleware/lib/infra/server"
	"os"
)

type InvokerImpl struct {}

func (InvokerImpl) Invoke() {
	s := server.NewServerRequestHandlerImpl("1234")

	//conexao := server.GetConn()
	fmt.Println("Início do for")

	//var req shared.Pessoa
	//msgFromClient := make([]byte, 1024)
	//var msgToClient []byte

	for {
		fmt.Println("1")

		msgToBeUnmarshalled := s.Receive()

		fmt.Println("chegou")

		msgReturned, err := Unmarshall(msgToBeUnmarshalled)

		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}


		/*erro3 = json.Unmarshal(msgFromClient[:n], &req)
		if erro3 != nil{
			fmt.Println(erro3)
			os.Exit(3)
		}*/

	}

	fmt.Println("Fim do loop")
}