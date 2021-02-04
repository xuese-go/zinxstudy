package ziface

import "net"

//定义连接
type IConnection interface {

	//	启动链接，让当前连接准备工作
	Start()
	//	停止连接，让当前链接停止工作
	Stop()
	//	获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	//	获取当前链接的连接ID
	GetConnID() uint32
	//	获取远程客户端Addr
	RemoteAddr() net.Addr
	//	发送数据
	Send([]byte) error
}

//处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
