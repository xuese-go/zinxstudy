package utils

import (
	"encoding/json"
	"github.com/xuese-go/zinxstudy/zinx/ziface"
	"io/ioutil"
)

/**
存储ZinX框架的全局参数
存储配置文件中的配置信息
*/

type GlobalObj struct {

	//	Server
	TcpServer ziface.IServer //当前zinx全局的Server对象
	Host      string         //服务器监听的IP
	Port      int            //服务器监听的端口
	Name      string         //服务器名称

	//	zinx
	Version        string //zinx版本
	MaxConn        int    //当前服务器主机允许的最大连接数
	MaxPackageSize uint32 //当前zinx数据包的最大值
}

//对外提供全局参数
var GlobalObject *GlobalObj

func (g *GlobalObj) loadJson() {
	data, err := ioutil.ReadFile("config/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObj{})
	if err != nil {
		panic(err)
	}

}

//初始化参数
func init() {
	GlobalObject = &GlobalObj{
		Name:           "zinx",
		Version:        "v0.4",
		Port:           9100,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//加载配置文件中的配置
	GlobalObject.loadJson()
}
