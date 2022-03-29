package znet

import "zinx/zInterface"

// Request 封装客户端发送的请求信息
type Request struct {
	conn zInterface.IConn
	data []byte
}

// GetReqConn 获取请求的连接信息
func (r Request) GetReqConn() zInterface.IConn {
	return r.conn
}

// GetReqData 获取请求数据
func (r Request) GetReqData() []byte {
	return r.data
}

func NewRequest(conn zInterface.IConn, data []byte) zInterface.IRequest {
	return &Request{
		conn,
		data,
	}
}
