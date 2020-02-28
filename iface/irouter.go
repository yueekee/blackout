package iface

type IRouter interface {
	PreHandle(request IRequest)			// 主业务之前处理的前置业务
	Handle(request IRequest)			// 处理conn业务的方法
	PostHandle(request IRequest)		// 主业务之后处理的后置业务
}

