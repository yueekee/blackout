package iface

type IServer interface {
	Start()
	Stop()
	Serve()			// 开启业务服务方法
}
