package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
	A, B int
}

type Result int

type RpcServer struct{}

func (t RpcServer) Add(args *Args, result *Result) error {
	log.Printf("Adding %d to %d\n", args.A, args.B)
	*result = Result(args.A + args.B)
	return nil
} 

const addr = ":7070"

func main() {
	go createServer(addr)
	time.Sleep(1 * time.Second)
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result

	err = client.Call("RpcServer.Add", args, &result)
	if err != nil {
		log.Fatalf("error in RpcServer", err)
	}
	log.Printf("%d+%d=%d\n", args.A, args.B, result)
}

func createServer(addr string) {
	server := rpc.NewServer()
	err := server.Register(RpcServer{})
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Couldn't start listening on %s errors: %s",
							 addr, err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}