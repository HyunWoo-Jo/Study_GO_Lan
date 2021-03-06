package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "222.111.17.41:8000")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096)
		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n]))

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "hello " + strconv.Itoa(i)

			c.Write([]byte(s))
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
