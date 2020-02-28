package main

import (
	"fmt"
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

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("write err:", err)
			return
		}
		buf := make([]byte, 51)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Printf("\tserver call back : #%s#\n", buf)
		fmt.Printf("\tserver call back : read = #%d#\n", read)
		time.Sleep(1 * time.Second)
	}
}