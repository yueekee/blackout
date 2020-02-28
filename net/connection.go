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
	handleAPI		iface.HandFunc
	ExitBuffChan	chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api iface.HandFunc) *Connection {
	return &Connection{
		Conn: 			conn,
		ConnID: 		connID,
		isClose: 		false,
		handleAPI: 		callback_api,
		ExitBuffChan: 	make(chan bool, 1),
	}
}

func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
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
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err:", err)
			c.ExitBuffChan <-true
			continue
		}

		if err := c.handleAPI(c.Conn, buf, n); err != nil {
			fmt.Println("connID", c.ConnID, "handle err")
			c.ExitBuffChan <-true
			return
		}
	}
}

func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}