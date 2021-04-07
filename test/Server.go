package main

import (
	"fmt"
	"github.com/liankui/blackout/iface"
	"github.com/liankui/blackout/inet"
)

type PingRouter struct {
	inet.BaseRouter
}

//func (this *PingRouter) PreHandle(request iface.IRequest) {
//	fmt.Println("Call Router PreHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping......\n"))
//	if err != nil {
//		fmt.Println("call back ping err:", err)
//	}
//}

func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle")

	// 读取客户端数据
	fmt.Printf("receive msg from client, msgID=%v, data=%v\n", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(1, []byte("ping......\n"))
	if err != nil {
		fmt.Println("call back ping err:", err)
	}
}

//func (this *PingRouter) PostHandle(request iface.IRequest) {
//	fmt.Println("Call Router PostHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping......\n"))
//	if err != nil {
//		fmt.Println("call back ping err:", err)
//	}
//}

func main() {
	server := inet.NewServer()
	server.AddRouter(&PingRouter{})
	server.Serve()
}
