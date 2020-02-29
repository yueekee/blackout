package main

import (
	"fmt"
	"github.com/yueekee/blackout/inet"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:7777")
	if e != nil {
		fmt.Println("client dial err:", e)
		return
	}
	fmt.Println("dial success")

	// 创建一个封包对象dp
	dp := inet.NewDataPack()
	// 封装一个msg1包
	msg1 := &inet.Message{
		ID: 		0,
		DataLen: 	5,
		Data:  		[]byte{'h','e','l','l','o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
		return
	}
	msg2 := &inet.Message{
		ID: 		1,
		DataLen: 	6,
		Data:  		[]byte{'w','o','r','l','d','!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err:", err)
		return
	}
	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)
	fmt.Println("数据已经发送，等待服务器响应")
	// 客户端阻塞
	select {}
}
