package rpcexe

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/HelloService"

type HelloServiceInterface interface {
	Hello(req string, reply *string) error
}

type HelloService struct{}

func (p HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func GenerateService() {

	service := RegisterHelloService(HelloService{})
	if service != nil {
		log.Fatal(service)
	}
	listener, e := net.Listen("tcp", ":15511")

	if e != nil {
		log.Fatal("listen tcp error:", e)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
