package main

import "zinx/znet"

/*
	基于Zinx的服务端
*/
func main() {
	// 1.创建一个Server
	server := znet.NewServer("[zinx v0.1]")
	// 2.启动服务器
	server.Serve()
}
