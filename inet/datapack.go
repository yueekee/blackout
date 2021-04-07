package inet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/liankui/blackout/iface"
	"github.com/liankui/blackout/utils"
)

type DataPack struct {}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

// 封包方法（压缩数据）
func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	// 存放byte字节的缓存
	dataBuf := bytes.NewBuffer([]byte{})

	// 写msgID
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// 写dataLen
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// 写data
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

// 拆包方法（解压数据）
func (dp *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {
	// 输入二进制数据的ioReader
	dataBuf := bytes.NewReader(binaryData)

	// 解压head信息，得到msgID和dataLen
	msg := &Message{}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 判断dataLen的长度是否超过允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too long msg data received")
	}

	return msg, nil
}
