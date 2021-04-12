package znet

import "github.com/xuese-go/zinxstudy/zinx/ziface"

type Request struct {

	//	建立好的链接
	conn ziface.IConnection

	//	链接中发送的数据
	data []byte
}

//	得到当前的链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//	得到当前链接中的数据
func (r *Request) GetData() []byte {
	return r.data
}
