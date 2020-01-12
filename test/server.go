package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

	defer conn.Close()
}
