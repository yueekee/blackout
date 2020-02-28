package net

import (
	"fmt"
	"github.com/yueekee/blackout/iface"
	"net"
	"time"
)

type Server struct {
	Name 		string
	IPVersion	string		// 网络，tcp/tcp4/tcp6
	IP			string
	Port		int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n",
		s.IP, s.Port)

	// 开启listener的业务
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err:", err)
			return
		}

		fmt.Println("start server ", s.Name, " success, now listening...")

		// 启动server网络连接服务
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// TODO Server.Start() 设置服务器的最大连接次数，若超则关闭新连接
			// TODO Server.Start() 处理新连接请求的业务方法，handler和conn绑定

			go func() {
				// 不断地循环从客户端获取数据
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf err ", err)
						continue
					}
					// 回显
					_, err = conn.Write(buf[:n])
					if err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Server , name " , s.Name)

	// TODO Server.Stop() 将其他需要清理的连接信息一并清理掉
}

func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve() 启动服务时需要做的其他事情

	for {
		time.Sleep(time.Second * 10)
	}
}

func NewServer(name string) iface.IServer {
	return &Server{
		Name:		name,
		IPVersion:	"tcp4",
		IP:			"0.0.0.0",
		Port:		7777,
	}
}
