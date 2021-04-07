package main

import (
	"fmt"
	"github.com/liankui/blackout/inet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, e := net.Dial("tcp", "127.0.0.1:7777")
	if e != nil {
		fmt.Println("Dial err:", e)
		return
	}
	defer conn.Close()

	for {
		dp := inet.NewDataPack()
		msg, _ := dp.Pack(inet.NewMsgPackage(0, []byte("-----Client test msg-----")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write err:", err)
			return
		}

		// 先读取流中head部分
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read err:", err)
			return
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

		time.Sleep(1 * time.Second)
	}
}
