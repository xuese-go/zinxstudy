package ziface

/*
将客户端的请求的链接信息以及数据进行封装
*/
type IRequest interface {
	//	得到当前的链接
	GetConnection() IConnection
	//	得到当前链接中的数据
	GetData() []byte
}
