package znet

import (
	"fmt"
	"net"
	"strings"
	"zinx/zInterface"
)

type Conn struct {
	Conn       *net.TCPConn          //当前连接的socket TCP套接字
	ID         uint32                // 当前连接的ID
	isClosed   bool                  // 当前连接的状态
	HandleFunc zInterface.HandleFunc // 当前连接绑定的处理业务的方法
	ExitChan   chan bool             // 告知当前连接该断开
	Router     zInterface.IRouter    // router
}

func NewConn(conn *net.TCPConn, id uint32, router zInterface.IRouter) *Conn {
	c := &Conn{
		Conn:     conn,
		ID:       id,
		isClosed: false,
		//HandleFunc: handleFunc,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// StartReader 连接的读业务方法
func (c *Conn) StartReader() {
	defer fmt.Println("Reader is exist..Remote addr is", c.GetRemoteAddr().String())
	defer c.Stop()
	conn := c.Conn
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			return
		}
		// 打印接受的信息
		fmt.Printf("received from %s:%s\n", conn.RemoteAddr().String(), strings.TrimSpace(string(buf[:n])))
		// 包装请求
		request := NewRequest(c, buf[:n])
		fmt.Println(c.Router == nil)
		go func() {
			c.Router.PreHandler(request)
			c.Router.Handler(request)
			c.Router.PostHandler(request)
		}()
	}

}

func (c *Conn) Start() {
	// 提示信息
	fmt.Printf("Conn Start..ID=%d\n", c.ID)
	// 启动读数据的业务
	go c.StartReader()
	//go c.StartWriter()
}

func (c *Conn) Stop() {
	// 提示信息
	fmt.Printf("Conn Stop..ID=%d\n", c.ID)

	if c.isClosed {
		return
	}
	c.isClosed = true
	// 释放连接
	c.Conn.Close()
	// 回收资源
	close(c.ExitChan)
}

func (c *Conn) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Conn) GetConnID() uint32 {
	return c.ID
}

func (c *Conn) GetRemoteAddr() net.Addr {
	return c.GetTCPConn().RemoteAddr()
}

func (c *Conn) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}
