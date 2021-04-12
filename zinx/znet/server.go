package znet

import (
	"errors"
	"fmt"
	"github.com/xuese-go/zinxstudy/zinx/ziface"
	"net"
)

//定义服务器
type Server struct {
	//	服务器名称
	Name string
	//	服务器监听的IP
	IP string
	//	IP版本
	IPVersion string
	//	服务器监听的端口
	Port int
}

func (s *Server) Start() {

	go func() {

		//	1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("tcpAddr err:", err)
			return
		}
		//	2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("tcpListener err:", err)
		}

		fmt.Println("Server Start success, ip:", s.IP, " ipVersion:", s.IPVersion, " port:", s.Port)
		var cid uint32 = 0
		//	3.阻塞的等待客户连接，处理客户端连接业务
		for true {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err:", err)
				continue
			}
			//	客户端已经与服务器建立连接，做最基本的回写最大为512字节的回写
			//	调用链接 绑定业务处理 回显数据
			dealConn := NewConnection(conn, cid, func(conn *net.TCPConn, data []byte, cnt int) error {
				//回显业务
				if _, err := conn.Write(data[:cnt]); err != nil {
					return errors.New("write to client err")
				}
				return nil
			})
			cid++

			//	启动链接
			go dealConn.Start()
		}
	}()
}

func (s *Server) Server() {
	s.Start()

	//TODO 阻塞后，此处可以做一些服务启动之后的额外业务

	//	阻塞
	select {}
}

func (s *Server) Stop() {

}

//初始化server
func NewServer(name string, port int) ziface.IServer {
	s := &Server{
		Name:      name,
		IP:        "localhost",
		IPVersion: "tcp4",
		Port:      port,
	}
	return s
}
