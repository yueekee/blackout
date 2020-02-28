package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn // 获取当前链接的socket
	GetConnID() uint32              // 当前链接ID
	RemoteAddr() net.Addr           // 远程客户端地址
}

// 定义一个统一处理 链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
