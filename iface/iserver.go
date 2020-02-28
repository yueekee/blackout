package iface

type IServer interface {
	Start()
	Stop()
	Serve()			// 开启业务服务方法
	AddRouter(router IRouter)	//给当前服务注册一个路由业务方法，供客户端链接处理使用
}
