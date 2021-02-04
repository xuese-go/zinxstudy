package main

import "github.com/xuese-go/zinxstudy/zinx/V0.1/znet"

/**
zinx 服务demo
*/
func main() {

	s := znet.NewServer("[zinx V0.1]", "localhost", 9000)

	s.Server()
}
