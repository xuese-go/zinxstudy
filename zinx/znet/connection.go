package znet

import (
	"fmt"
	"github.com/xuese-go/zinxstudy/zinx/ziface"
	"net"
)

type Connection struct {
	//链接套接字
	Conn *net.TCPConn
	//	链接ID
	ConnID uint32
	//	链接状态
	isClosed bool
	//告知当前连接已停止/退出 channel
	ExitChan chan bool
	//	该链接处理的Router
	Router ziface.IRouter
}

//初始化链接模块
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		isClosed: false,
		Router:   router,
		//1字节缓冲的 bool类型通道
		ExitChan: make(chan bool, 1),
	}
	return c
}

//读取数据
func (c *Connection) StartRead() {
	fmt.Println("Connection Read start...")
	defer fmt.Println("conn read is close,connId=", c.ConnID)
	defer c.Stop()
	for true {
		//读取数据
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read err", err)
			continue
		}
		fmt.Println("接收到的数据：", buf)

		//得到当前Conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		//	从当前路由中调用绑定的Conn对应的Router
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

//	启动链接，让当前连接准备工作
func (c *Connection) Start() {
	fmt.Println("Connection Start...")
	go c.StartRead()
}

//	停止连接，让当前链接停止工作
func (c *Connection) Stop() {
	if !c.isClosed {
		c.isClosed = true
	}
	//	关闭socket
	_ = c.Conn.Close()
	//	关闭管道,回收资源
	close(c.ExitChan)
}

//	获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//	获取当前链接的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//	获取远程客户端Addr
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//	发送数据
func (c *Connection) Send([]byte) error {
	return nil
}
