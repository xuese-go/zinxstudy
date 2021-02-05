package main

import "github.com/xuese-go/zinxstudy/zinx/V0.2/znet"

/**
zinx 服务demo
*/
func main() {

	s := znet.NewServer("[zinx V0.2]", "localhost", 9000)

	s.Server()
}
