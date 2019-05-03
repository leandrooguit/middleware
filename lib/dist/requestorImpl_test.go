package dist

import (
	"reflect"
	"testing"
)

func TestRequestorImpl_Invoke(t *testing.T) {
	type args struct {
		inv InvocationImpl
	}
	tests := []struct {
		name string
		r    RequestorImpl
		args args
		want Termination
	}{
		{"Teste 1",
			RequestorImpl{},
			args{InvocationImpl{1000, "127.0.0.1", 1234, "play", [2]string{"P", "T"}}},
			TerminationImpl{},
		},
		{"Teste 2",
			RequestorImpl{},
			args{InvocationImpl{1000, "127.0.0.1", 1234, "play", [2]string{"P", "P"}}},
			TerminationImpl{},
		},
		{"Teste 3",
			RequestorImpl{},
			args{InvocationImpl{1000, "127.0.0.1", 1234, "play", [2]string{"T", "P"}}},
			TerminationImpl{},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := tt.r.Invoke(tt.args.inv)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestorImpl.Invoke() = %v, want %v", got, tt.want)
			}

		})

	}
}

/*func TestInvoker(t *testing.T) {
	r := new(RequestorImpl)

	var parameters [2]string
	parameters[0] = "Pedra"
	parameters[1] = "Tesoura"

	i := InvocationImpl{ObjectId: 1000, IpAddress: "127.0.0.1:1234", PortNumber: 1234, OperationName: "jankenpo", Parameters: parameters}

	r.Invoke(i)
}*/
