package ziface

type IDataPack interface {
	//	获取包头长度
	GetHeadLen() uint32
	//	封包方法
	Pack(msg IMessage) ([]byte, error)
	//	拆包方法
	Unpack(data []byte) (msg IMessage, err error)
}
