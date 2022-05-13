package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播Chan
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return &server
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]:" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message方法,用来进行消息广播
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		// 将msg发送给全部的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	// 当前业务
	fmt.Println("连接建立成功")

	user := NewUser(conn, this)
	user.Online()

	// 监听用户活跃
	isLive := make(chan bool)

	// 接受客户端消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil {
				fmt.Println("Conn Read Error")
				return
			}
			// 提取字符串，并取出\n
			msg := string(buf[:n-1])
			// 将消息进行广播
			user.DoMessage(msg)
			// 用户操作
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，重置下面的计时器
		case <-time.After(time.Second * 100):
			// 已经超时，强制踢出
			user.SendMsg("你被踢了...")
			// 销毁资源
			close(user.C)
			conn.Close()
			return
		}
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	// 1. socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 2. close listen socket
	defer listener.Close()

	// 启动监听
	go this.ListenMessage()
	for {
		// 3. accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accept err:", err)
			continue
		}

		// 4. do handler
		go this.Handler(conn)
	}
}
