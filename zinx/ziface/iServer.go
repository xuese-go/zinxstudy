package ziface

//定义服务器接口
type IServer interface {

	//	启动方法
	Start()
	//	运行方法
	Server()
	//	停止方法
	Stop()
	//	添加路由
	AddRouter(router IRouter)
}
