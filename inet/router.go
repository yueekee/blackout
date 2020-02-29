package inet

import "github.com/yueekee/blackout/iface"

type BaseRouter struct {}

func (br *BaseRouter) PreHandle(req iface.IRequest)  {}
func (br *BaseRouter) Handle(req iface.IRequest)     {}
func (br *BaseRouter) PostHandle(req iface.IRequest) {}
