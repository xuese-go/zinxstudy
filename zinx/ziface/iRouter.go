package ziface

/*
路由
*/

type IRouter interface {
	//	处理业务之前的钩子
	PreHandle(request IRequest)
	//	处理业务的主方法
	Handle(request IRequest)
	//	处理业务之后的钩子
	PostHandle(request IRequest)
}
