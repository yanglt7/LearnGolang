package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(filename string, conn net.Conn) {
	//新建文件
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}

	//写入文件
	buf := make([]byte, 1024*4)
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err: ", err1)
			}
			return
		}
		f.Write(buf[:n])
	}

	defer f.Close()
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	//建立连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err: ", err1)
		return
	}

	defer conn.Close()

	//读取文件名
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn Read err: ", err2)
		return
	}

	filename := string(buf[:n])

	//回复"ok"
	conn.Write([]byte("ok"))

	//接收文件
	RecvFile(filename, conn)
}
