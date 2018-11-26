package native

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/profilux-go/protocol"

	"github.com/cjburchell/yasls-client-go"
)

type nativeProtocol struct {
	Connection *connection
	Address    int
}

var errProtocol = errors.New("protocol error")

func codeError(code byte) error {
	return errors.WithStack(fmt.Errorf("code error %d", code))
}

// NewProtocol creates a http protocol
func NewProtocol(address string, port int, controller int) (protocol.IProtocol, error) {
	con, err := newConnection(address, port)
	if err != nil {
		return nil, err
	}

	var p nativeProtocol

	p.Address = controller
	p.Connection = con

	return &p, nil
}

func (protocol nativeProtocol) Disconnect() {
	protocol.Connection.Disconnect()
}

func (protocol nativeProtocol) SendData(code, data int) error {
	sendCommandInt(code, data, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return err
		}

		err = verifyAckPacket(reply)
		if err == nil {
			return nil
		}

		if err != errProtocol {
			return err
		}
	}
}

func (protocol nativeProtocol) SendText(code int, data string) error {
	sendTextCommand(code, data, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return err
		}

		err = verifyAckPacket(reply)
		if err == nil {
			return nil
		}

		if err != errProtocol {
			return err
		}
	}
}

func (protocol nativeProtocol) GetDataText(code int) (string, error) {
	sendCommand(code, protocol.Connection, protocol.Address)
	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return "", err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			return getMessageString(reply), nil
		}

		if err != errProtocol {
			return "", err
		}
	}
}

func (protocol nativeProtocol) GetData(code int) (int, error) {

	log.Printf("Sending GetData %d", code)
	sendCommand(code, protocol.Connection, protocol.Address)
	for {

		log.Printf("Reading Repsonce")
		reply, err := protocol.readPacket()
		if err != nil {
			return 0, err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			log.Printf("Got valid data")
			return getMessageData(reply), nil
		}

		if err != errProtocol {
			return 0, err
		}
	}
}

func (protocol nativeProtocol) GetDataShortArray(code int) ([]int, error) {
	sendCommand(code, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return nil, err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			return getMessageDataShortArray(reply), nil
		}

		if err != errProtocol {
			return nil, err
		}
	}
}

func (protocol nativeProtocol) GetDataByteArray(code int) ([]byte, error) {
	sendCommand(code, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return nil, err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			return getMessageBytes(reply), nil
		}

		if err != errProtocol {
			return nil, err
		}
	}
}

func (protocol nativeProtocol) GetDataTwoByteArray(code int) ([]int, error) {
	sendCommand(code, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return nil, err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			return getMessageDataTwoByteArray(reply), nil
		}

		if err != errProtocol {
			return nil, err
		}
	}
}

func (protocol nativeProtocol) GetDataBoolArray(code int) ([]bool, error) {
	sendCommand(code, protocol.Connection, protocol.Address)

	for {
		reply, err := protocol.readPacket()
		if err != nil {
			return nil, err
		}

		err = verifyDataPacket(reply, code)
		if err == nil {
			return getMessageBools(reply), nil
		}

		if err != errProtocol {
			return nil, err
		}
	}
}

func (protocol nativeProtocol) readPacket() (reply []byte, err error) {
	for {
		data, size, err := protocol.Connection.Read(1)
		if err != nil {
			return nil, err
		}

		if size == 0 {
			break
		}

		reply = append(reply, data...)

		if atEndOfPacket(reply) {
			break
		}
	}

	return
}

func verifyDataPacket(reply []byte, code int) error {
	if len(reply) < 4 {
		// strange packet size!
		log.Warn("Expecting Packet size of at least 4")
		return errProtocol
	}

	if reply[4] == eot {
		log.Warn("Unexpected Message: Empty Reply")
		return errProtocol
	}

	if reply[4] == stx {
		if reply[5] == nak {
			errorCode := reply[6]
			return codeError(errorCode)
		}

		if reply[5] == ack {
			log.Warn("Unexpected Message: ack")
			return errProtocol
		}

		// should be ok we must now look for the code and verify it
		replyCode := getGetMessageCode(reply)
		if replyCode != code {
			log.Warnf("Unexpected Message: Wrong Code Expecting %d Got %d", code, replyCode)
			return errProtocol
		}
	} else {
		log.Warn("Unknown message type")
		return errProtocol
	}

	return nil
}

func verifyAckPacket(reply []byte) error {
	if len(reply) < 4 {
		// strange packet size!
		log.Warnf("Expecting Packet size of at least 4")
		return errProtocol
	}

	if reply[4] == eot {
		log.Warn("Unexpected Message: Empty Reply")
		return errProtocol
	}

	if reply[4] == stx {
		if reply[5] == nak {
			errorCode := reply[6]
			return codeError(errorCode)
		}

		if reply[5] != ack {
			replyCode := getGetMessageCode(reply)
			log.Warnf("Unexpected Message: Code %d", replyCode)
			return errProtocol
		}

	} else {
		log.Warn("Unknown message type")
		return errProtocol
	}

	return nil
}
