package rpcexe

import (
	"log"
	"net/rpc"
)


func GenerateClient() {
	//client, e := rpc.Dial("tcp", "localhost:15511")
	//if e != nil {
	//	log.Fatal(e)
	//}
	//
	//var reply string
	//
	//call := client.Call(HelloServiceName+".Hello", "hello", &reply)
	//if call != nil {
	//	log.Fatal(call)
	//}
	//fmt.Println(reply)

	client, err := DialHelloService("tcp", "localhost:15511")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

}

type HelloServiceClient struct {
	*rpc.Client
}

//var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
