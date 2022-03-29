package zInterface

type (
	// IRouter 路由模块接口
	IRouter interface {
		PreHandler(request IRequest)  // 处理业务之前的Handler
		Handler(request IRequest)     // 处理业务的Handler
		PostHandler(request IRequest) // 处理业务之后的Handler
	}
)
