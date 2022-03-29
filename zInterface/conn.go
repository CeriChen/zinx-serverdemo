package zInterface

import "net"

//封装连接

type IConn interface {
	Start()                   // 启动连接
	Stop()                    // 关闭连接
	GetTCPConn() *net.TCPConn // 获取绑定的socket connection
	GetConnID() uint32        // 获取连接ID
	GetRemoteAddr() net.Addr  // 获取远程客户端Addr
	Send(data []byte) error   // 发送数据到远程客户端
}

// HandleFunc 处理业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
