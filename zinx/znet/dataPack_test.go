package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

/**
封包/拆包测试
*/
func TestDataPack(t *testing.T) {

	//1、创建socketTCP
	serverTest()
	//2、从客户端读取数据，拆包处理
	clientTest()
}

/**
模拟服务器
*/
func serverTest() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("server listen err:", err)
	}

	go func() {
		//接收链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
		}

		//处理客户端请求
		go func(conn net.Conn) {
			//	拆包过程
			dp := NewDataPack()
			for {
				//	第一步读取head，获取数据长度
				//声明用于存放头部数据的数组
				headData := make([]byte, dp.GetHeadLen())
				//读满指定长度的数据
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read data err:", err)
					break
				}
				//进行拆包
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("unpack head err:", err)
					break
				}
				//	判断头部长度确定是否由数据
				if msgHead.GetDataLen() > 0 {
					//	第二部根据数据长度获取data
					//	类型断言，类似java的强转
					msg := msgHead.(*Message)
					msg.Data = make([]byte, msgHead.GetDataLen())
					//之前把头部读取出来了，根据字节流读一节少一节的特性，所以此处不需要偏移了
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("read data err:", err)
						return
					}
					//输出读取的数据
					//把二进制数据转成string输出
					fmt.Println("READ MSG ID:", msg.Id, "MSG DATA:", string(msg.Data))
				} else {
					fmt.Println("data head len is 0")
				}
			}
		}(conn)
	}()

}

/**
模拟客户端
*/
func clientTest() {
	//Dial 拨号
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client dial err:", err)
	}

	//	创建一个封包对象
	dp := NewDataPack()
	//模拟拆包过程，创建两个一起发送
	//封装第一个包
	s1 := "Hello Word!"
	msg1 := &Message{
		Id:      1,
		DataLen: uint32(len(s1)),
		Data:    []byte(s1),
	}
	msgData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack data 1 err:", err)
		return
	}
	//封装第二个包
	s2 := "This is go TCP!"
	msg2 := &Message{
		Id:      2,
		DataLen: uint32(len(s2)),
		Data:    []byte(s2),
	}
	msgData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack data 2 err:", err)
		return
	}
	//将两个包黏在一起
	//不带...将会变成二维数组
	msgData1 = append(msgData1, msgData2...)
	//一次性发送两个粘在一起的大包
	if _, err := conn.Write(msgData1); err != nil {
		fmt.Println("client send err:", err)
	}

	//客户端阻塞
	time.Sleep(5000)
}
