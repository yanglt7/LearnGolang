package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读形式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err: ", err1)
			}
			return
		}

		conn.Write(buf[:n])
	}
}

func main() {
	//提示输入文件
	fmt.Println("请输入文件名： ")
	var path string
	fmt.Scan(&path)

	info, err3 := os.Stat(path)
	if err3 != nil {
		fmt.Println("os.Stat err: ", err3)
		return
	}

	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer conn.Close()

	//发送文件名
	_, err1 := conn.Write([]byte(info.Name()))
	if err1 != nil {
		fmt.Println("conn.Write err: ", err1)
		return
	}

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err: ", err2)
		return
	}

	fmt.Println("n: ", n)
	if "ok" == string(buf[:n]) {
		SendFile(path, conn)
	}

}
