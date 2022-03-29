package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client starting..")
	addr, err := net.ResolveTCPAddr("tcp4", ":8999")
	if err != nil {
		fmt.Println(addr)
		return
	}
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	// 读数据
	go func() {
		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(strings.TrimSpace(string(buf[:n])))
		}

	}()
	// 写数据并且阻塞
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = conn.Write([]byte(s))
		if err != nil {
			return
		}
	}
}
