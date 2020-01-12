package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client
var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息前会阻塞

		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func HandleConn(conn net.Conn) {
	cliAddr := conn.RemoteAddr().String()

	//创建结构体，默认用户名为网络地址
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线
	message <- MakeMsg(cli, "login")
	//提示我是谁
	cli.C <- MakeMsg(cli, "I am Here")

	isQuit := make(chan bool)
	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list: \n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				message <- MakeMsg(cli, msg)
			}

			hasData <- true
		}
	}()

	for {
		//通过select检测channel流动，检测超时
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr) // 当前用户从map移除
			message <- MakeMsg(cli, "logout")
			return
		case <-hasData:
		case <-time.After(30 * time.Second):
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "time out leave out")
			return
		}
	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，给每个成员发送消息
	go Manager()

	//主协程，阻塞监听用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConn(conn) //处理用户连接
	}

}
