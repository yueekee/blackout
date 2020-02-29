package main

import (
	"fmt"
	"github.com/yueekee/blackout/inet"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("建立了listener")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
		}

		// 处理客户端请求
		go func(conn net.Conn) {
			defer conn.Close()
			// 创建拆包封包对象dp
			dp := inet.NewDataPack()

			for {
				// 1 先读出流中head部分
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)		// 从conn中读取字节数据到headData中
				if err != nil {
					fmt.Println("read head error")
					break
				}

				// 将headData字节流拆包到msg中
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err:", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					// msg是有data数据，需要再次读取data数据
					msg := msgHead.(*inet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack data err:", err)
						return
					}

					fmt.Println("==> Recv Msg: ID=", msg.ID,
						", len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
