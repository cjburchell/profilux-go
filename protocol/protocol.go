package protocol

// IProtocol interface for the protocol
type IProtocol interface {
	SendData(code, data int) error
	GetDataText(code int) (string, error)
	GetData(code int) (int, error)
	Disconnect()
}
