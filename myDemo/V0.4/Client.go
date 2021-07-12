package main

import (
	"fmt"
	"net"
	"time"
)

/**
模拟客户端
*/
func main() {

	//	连接服务器，得到conn
	conn, err := net.Dial("tcp", "localhost:9100")
	if err != nil {
		fmt.Println("dial err:", err)
	}

	for true {
		//	发送消息
		_, err = conn.Write([]byte("Hello Word!3.0\r"))
		if err != nil {
			fmt.Println("conn write err:", err)
		}

		//	读取消息
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err:", err)

		}

		fmt.Printf("server call back msg:%s", buf)

		//	阻塞
		time.Sleep(time.Second * 1)
	}

}
