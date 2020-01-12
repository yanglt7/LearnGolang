package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	buf := make([]byte, 2048)
	addr := conn.RemoteAddr().String()
	fmt.Println("connect success.")

	defer conn.Close()
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s] -> %s\n", addr, string(buf[:n]))
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, " exit")
			return
		}

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
		}
		go HandleConn(conn)
	}

}
