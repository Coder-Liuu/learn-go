package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端
	client := Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	// 进行连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("Conn error")
		return nil
	}
	client.conn = conn
	return &client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.改名字")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法的命令...")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>> 请输入你的新用户名字:")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"

	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	var chatMsg string
	// 提示用户输入
	fmt.Println(">>> 请输入你要发送的内容(输入exit退出)：")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			chatMsg = chatMsg + "\n"
			_, err := client.conn.Write([]byte(chatMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>> 请输入你要发送的内容(输入exit退出)：")
		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write error!")
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>请输入聊天用户名字(exit退出)")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容(exit退出)")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				chatMsg = "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(chatMsg))
				if err != nil {
					fmt.Println("conn.Write err:", err)
					break
				}
			}
			fmt.Println(">>>>请输入消息内容(exit退出)")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println(">>>>请输入聊天用户名字(exit退出)")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
		case 2:
			// 私聊模式
			client.PrivateChat()
		case 3:
			// 更新用户名字
			client.UpdateName()
		}
	}
}

// 处理服务器的消息
func (client *Client) DealResponse() {
	// 一旦client.conn有消息就直接拷贝到stdout中，永久阻塞
	io.Copy(os.Stdout, client.conn)
}

var serverIp string
var serverPort int

// client -ip Ip地址 -port 端口值
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认8888)")
}

func main() {
	// 解析命令行
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>> 连接服务器失败...")
		return
	}
	fmt.Println(">>> 连接服务器成功...")

	go client.DealResponse()
	client.Run()
}
