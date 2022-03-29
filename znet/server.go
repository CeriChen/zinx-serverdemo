package znet

import (
	"fmt"
	"net"
	"zinx/zInterface"
)

// Server IServer的接口实现
type Server struct {
	Name      string             // 服务器名称
	IPVersion string             // 服务器绑定的IP版本
	IP        string             // 服务器监听的IP
	Port      int                // 服务器监听的端口
	Router    zInterface.IRouter // 服务器路由
}

// NewServer 初始化Server
func NewServer(name string) zInterface.IServer {
	newServer := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return newServer
}

// AddRouter 添加路由
func (s *Server) AddRouter(r zInterface.IRouter) {
	fmt.Println("add router success!")
	s.Router = r
}

// CallBackAPI 回显信息到客户端--以后可以自定义业务处理API
//func CallBackAPI(conn *net.TCPConn, data []byte, n int) error {
//	fmt.Println("[Conn Handle] CallbackToClient..")
//	if _, err := conn.Write(data[:n]); err != nil {
//		return errors.New("call back to client failed error")
//	}
//	return nil
//}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener At IP: %s, Port: %d\n", s.IP, s.Port)
	go func() {
		// 1.获取TCP的Addr(Socket句柄)
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		// 2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err:", err)
			return
		}
		fmt.Println("Start Zinx Server success.", s.Name, " listening...")
		// 分配cid
		var cid uint32 = 0
		// 3.等待客户端连接，处理请求
		for {
			// 如果没有客户端连接，阻塞
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept conn shutdown, err:", err)
				continue
			}
			// 封装连接并且处理连接

			c := NewConn(conn, cid, s.Router)
			c.Start()

			// 每次接受一个连接后
			cid++
		}
	}()
}

// Stop 停止服务器
func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// TODO 启动服务器之后的一些业务

	// 阻塞
	select {}
}
