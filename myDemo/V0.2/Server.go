package main

import "github.com/xuese-go/zinxstudy/zinx/znet"

/**
zinx 服务demo
*/
func main() {

	s := znet.NewServer("[zinx V0.2]", 9100)

	s.Server()
}
