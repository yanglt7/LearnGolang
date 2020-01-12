package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("hello world."))
}
