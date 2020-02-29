package inet

type Message struct {
	ID 			uint32
	DataLen 	uint32
	Data 		[]byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID: 		id,
		DataLen: 	uint32(len(data)),
		Data: 		data,
	}
}

func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetMsgID(id uint32) {
	msg.ID = id
}

func (msg *Message) SetDataLen (len uint32) {
	msg.DataLen = len
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}