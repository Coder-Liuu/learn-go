package main

import (
	_ "fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string // channal 通道
	conn   net.Conn    // 连接
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听
	go user.ListenMessage()
	return &user
}

// 用户上线
func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "用户进行上线")
}

// 用户下线
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "用户下线")
}

// 给当前用户发送消息 API
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户发送消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询所有在线用户信息
		this.SendMsg("---查询结果---\n")
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			res := user.Addr + ":" + user.Name + ": 在线...\n"
			this.SendMsg(res)
		}
		this.server.mapLock.Unlock()
		this.SendMsg("--------------\n")
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式rename|新名字
		newName := msg[7:]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前名称已存在，用户名更新失败!\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("你已经更新用户名为:" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式to|张三|内容
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用<to|张三|内容>的格式\n")
			return
		}
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("当前用户不存在\n")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息，请重新发送\n")
			return
		}

		content = remoteUser.Name + "悄悄的对你说:" + content + "\n"
		remoteUser.SendMsg(content)
		this.SendMsg("消息发送完毕\n")
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前用户的 channal,一旦有消息就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.SendMsg(msg + "\n")
	}
}
