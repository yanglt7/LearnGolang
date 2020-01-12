package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer conn.Close()

	go func() {
		str := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err: ", err)
				return
			}
			conn.Write(str[:n])
		}
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
