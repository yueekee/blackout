package inet

import (
	"errors"
	"fmt"
	"github.com/yueekee/blackout/iface"
	"io"
	"net"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       uint32
	isClose      bool
	router       iface.IRouter
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClose:      false,
		router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) Stop() {
	if c.isClose == true {
		return
	}
	c.isClose = true

	// 关闭socket链接
	c.Conn.Close()

	// 阻塞缓存队列读数据的业务
	c.ExitBuffChan <- true
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
		// 创建拆包解包对象dp
		dp := NewDataPack()

		// 读取客户端head的msg
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("receive head msg err:", err)
			c.ExitBuffChan <- true
			continue
		}

		// 拆包，得到msg（包含msgID和dataLen）
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err:", err)
			c.ExitBuffChan <- true
			continue
		}

		// 根据dataLen读取data至msg.data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data err:", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		// 得到当前客户端请求的Request数据
		req := Request{
			conn: 	c,
			msg:	msg,
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

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClose == true {
		return errors.New("Connection closed when send msg")
	}

	// 	将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Println("pack err, msgID:", msgID)
		return errors.New("Pack msg error")
	}

	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("write err, msgID:", msgID)
		c.ExitBuffChan <- true
		return errors.New("Conn write err")
	}

	return nil
}