package net

import (
	"errors"
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
	Router 		iface.IRouter
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

		var connID uint32 = 0
		// 启动server网络连接服务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// TODO Server.Start() 设置服务器的最大连接次数，若超则关闭新连接

			// 处理新连接请求的业务方法，将handler和conn绑定
			connection := NewConnection(conn, connID, s.Router)
			connID++
			// 启动当前链接的处理业务
			go connection.Start()
		}
	}()
}

// 服务器回显的封装api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
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

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
	fmt.Println("Add Router success!" )
}

func NewServer(name string) iface.IServer {
	return &Server{
		Name:		name,
		IPVersion:	"tcp4",
		IP:			"0.0.0.0",
		Port:		7777,
		Router:  	nil,
	}
}
