package znet

import "github.com/xuese-go/zinxstudy/zinx/ziface"

/**
实现router时，先嵌入BaseRouter基类，然后根据需要对这个类的方法进行重写。
实现重写需要实现所有的方法，但是实际业务中可能不需要全部方法，这里重写所有的方法，方便业务中根据需要只需要重写需要的方法即可。
*/
type BaseRouter struct{}

//	处理业务之前的钩子
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//	处理业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//	处理业务之后的钩子
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
