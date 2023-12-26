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

func (ms *MathService) Divide(args Args, reply *float64) error {
	if args.B == 0 {
		return fmt.Errorf("Can't divide by 0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)

	listener, err := net.Listen("tcp", ":1237")
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
