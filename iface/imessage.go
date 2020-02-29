package iface

type IMessage interface {
	GetMsgID()			uint32
	GetData()			[]byte			// 获取消息内容
	GetDataLen()		uint32

	SetMsgID(uint32)
	SetData([]byte)						// 设置消息内容
	SetDataLen(uint32)
}


