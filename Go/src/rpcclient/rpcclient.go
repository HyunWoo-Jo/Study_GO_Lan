package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {
	client, err := rpc.Dial("tcp", "222.111.17.41:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	args := &Args{1, 2}
	reply := new(Reply)
	err = client.Call("Call.Sum", args, reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil)
	<-sumCall.Done
	fmt.Println(reply.C)
}
