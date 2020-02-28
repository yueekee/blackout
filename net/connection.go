package net

import (
	"fmt"
	"github.com/yueekee/blackout/iface"
	"net"
)

type Connection struct {
	Conn			*net.TCPConn
	ConnID			uint32
	isClose			bool
	router 			iface.IRouter
	ExitBuffChan	chan bool
}


func (c *Connection) Stop() {
	if c.isClose == true {
		return
	}
	c.isClose = true

	// 关闭socket链接
	c.Conn.Close()

	// 阻塞缓存队列读数据的业务
	c.ExitBuffChan <-true
	// 关闭该链接全部通道
	close(c.ExitBuffChan)
}

func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			// 得到退出消息，不再阻塞
			return
		}
	}
}

// 处理conn读数据的goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err:", err)
			c.ExitBuffChan <-true
			continue
		}

		// 得到当前客户端请求的Request数据
		req := Request{
			conn:	c,
			data:	buf,
		}
		// 从路由中找到注册绑定conn对应的Handle
		go func(request iface.IRequest) {
			// 执行注册的路由方法
			c.router.PreHandle(request)
			c.router.Handle(request)
			c.router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClose:      false,
		router:		  router,
		ExitBuffChan: make(chan bool, 1),
	}
}