package zInterface

type IRequest interface {
	GetReqConn() IConn  // 获取请求的连接
	GetReqData() []byte // 获取请求数据
}
