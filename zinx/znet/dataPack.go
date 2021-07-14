package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/xuese-go/zinxstudy/zinx/utils"
	"github.com/xuese-go/zinxstudy/zinx/ziface"
)

/**
封包   拆包
*/
type DataPack struct {
}

//	获取包头长度
func (d *DataPack) GetHeadLen() uint32 {
	//头部由id和数据长度两部分组成，每部分长度4字节
	return 8
}

//	封包方法
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	//	将dataLen写入dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		fmt.Println("pack err:", err)
		return nil, err
	}
	//	将MsgId写入dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataId()); err != nil {
		fmt.Println("pack err:", err)
		return nil, err
	}
	//	将data数据写入dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		fmt.Println("pack err:", err)
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//	拆包方法
func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(data)
	msg := &Message{}

	//读取头部获取数据长度
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		fmt.Println("unpack err:", err)
		return nil, err
	}

	//读取头部ID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		fmt.Println("unpack err:", err)
		return nil, err
	}

	//	判断dataLen是否超出配置的值
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}

	return msg, nil
}

func NewDataPack() DataPack {
	return DataPack{}
}
