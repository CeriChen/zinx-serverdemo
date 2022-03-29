package main

import (
	"zinx/zInterface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PreHandler(req zInterface.IRequest) {
	req.GetReqConn().GetTCPConn().Write([]byte("before doing..."))
}
func (r *PingRouter) Handler(req zInterface.IRequest) {
	req.GetReqConn().GetTCPConn().Write(append([]byte("callback:"), req.GetReqData()...))
}
func (r *PingRouter) PostHandler(req zInterface.IRequest) {
	req.GetReqConn().GetTCPConn().Write([]byte("after..."))
}

func main() {
	server := znet.NewServer("zinx v0.3")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
