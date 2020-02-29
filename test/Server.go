package main

import (
	"fmt"
	"github.com/yueekee/blackout/iface"
	"github.com/yueekee/blackout/inet"
)

type PingRouter struct {
	inet.BaseRouter
}

func (this *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping......\n"))
	if err != nil {
		fmt.Println("call back ping err:", err)
	}
}

func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping......\n"))
	if err != nil {
		fmt.Println("call back ping err:", err)
	}
}

func (this *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping......\n"))
	if err != nil {
		fmt.Println("call back ping err:", err)
	}
}

func main() {
	server := inet.NewServer()
	server.AddRouter(&PingRouter{})
	server.Serve()
}
