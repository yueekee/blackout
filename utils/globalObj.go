package utils

import (
	"encoding/json"
	"github.com/yueekee/blackout/iface"
	"io/ioutil"
)

type GlobalObj struct {
	Name          string
	Version       string
	Host          string
	TcpPort       int
	TcpServer     iface.IServer
	MaxConn       int    // 服务器允许的最大连接个数
	MaxPacketSize uint32 // 所需数据包的最大值
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:          "APP1",
		Version:       "1.0.3",
		Host:          "0.0.0.0",
		TcpPort:       7777,
		MaxConn:       3,
		MaxPacketSize: 10000,
	}

	GlobalObject.Reload()
}

// 读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("json :%s\n", data)
	e := json.Unmarshal(data, &GlobalObject)
	if e != nil {
		panic(e)
	}
}
