package znet

import "zinx/zInterface"

type BaseRouter struct{}

func (r *BaseRouter) PreHandler(request zInterface.IRequest) {
}

func (r *BaseRouter) Handler(request zInterface.IRequest) {
}

func (r *BaseRouter) PostHandler(request zInterface.IRequest) {
}
