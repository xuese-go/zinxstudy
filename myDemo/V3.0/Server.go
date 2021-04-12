package main

import (
	"github.com/xuese-go/zinxstudy/zinx/ziface"
	"github.com/xuese-go/zinxstudy/zinx/znet"
)

/**
zinx 服务demo
*/

type PingRouter struct {
	znet.BaseRouter
}

func (br *PingRouter) PreHandle(request ziface.IRequest) {
	_, _ = request.GetConnection().GetTCPConnection().Write([]byte("Hello Word 3.0\r"))
}

func main() {

	s := znet.NewServer("[zinx V0.2]", 9100)

	s.AddRouter(&PingRouter{})

	s.Server()
}
