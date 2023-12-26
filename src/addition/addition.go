package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	A, B float64
}

type MathService struct{}

func (ms *MathService) Add(args Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
