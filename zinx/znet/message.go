package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

//	获取消息ID
func (m *Message) GetDataId() uint32 {
	return m.Id
}

//	获取消息长度
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

//	获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}

//	设置消息ID
func (m *Message) SetDataId(id uint32) {
	m.SetDataId(id)
}

//	设置消息长度
func (m *Message) SetDataLen(dataLen uint32) {
	m.SetDataLen(dataLen)
}

//	设置消息内容
func (m *Message) SetData(data []byte) {
	m.SetData(data)
}
