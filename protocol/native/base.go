package native

const (
	soh            byte = 0x01
	stx            byte = 0x02
	enq            byte = 0x05
	etx            byte = 0x03
	eot            byte = 0x04
	ack            byte = 0x06
	nak            byte = 0x15
	dataOffset     byte = 0x30
	codeOffset     byte = 0x40
	checksumOffset byte = 0x20
)

func calculateChecksum(data []byte) byte {
	dataSum := 0
	for b := range data {
		dataSum += b
	}

	bca := byte(dataSum % 256)
	if bca < checksumOffset {
		bca = bca + checksumOffset
	}

	return bca
}

func addData(command []byte, data int) []byte {
	if data == 0 {
		command = append(command, dataOffset) // the value of the data is zero we then add 0
	} else {
		for data != 0 {
			value := byte(data & 0xf)
			data >>= 4
			command = append(command, dataOffset|value)
		}
	}

	return command
}

func addTextData(command []byte, data string) []byte {
	for c := range data {
		value1 := byte(c & 0xf)
		value2 := byte(c >> 4)
		command = append(command, dataOffset|value1)
		command = append(command, dataOffset|value2)
	}
	return command
}

func addCode(command []byte, code int) []byte {
	for code != 0 {
		value := byte(code & 0xf)
		code >>= 4
		command = append(command, codeOffset|value)
	}

	return command
}

func sendEnd(connection *connection, address int) {
	command := []byte{soh, (byte)(0x50 + address), 0x70}
	command = append(command, calculateChecksum(command)) // BCA
	command = append(command, eot)                        // eot
	command = append(command, 0xFF)                       // eot
	command = append(command, 0xFF)                       // eot

	connection.Write(command)
}

func sendCommand(code int, connection *connection, address int) {
	var command = []byte{soh, (byte)(0x50 + address), 0x70}
	command = append(command, calculateChecksum(command)) // BCA
	command = append(command, stx)                        // stx
	command = addCode(command, code)                      // code
	command = append(command, enq)                        // enq
	command = append(command, etx)                        // etx
	command = append(command, calculateChecksum(command)) // BCC
	command = append(command, eot)                        // eot
	command = append(command, 0xFF)                       // eot
	command = append(command, 0xFF)                       // eot

	connection.Write(command)

	sendEnd(connection, address)
}

func sendCommandInt(code, data int, connection *connection, address int) {
	var command = []byte{soh, (byte)(0x50 + address), 0x70}
	command = append(command, calculateChecksum(command)) // BCA
	command = append(command, stx)                        // stx
	command = addCode(command, code)                      // code
	command = addData(command, data)                      // data
	command = append(command, enq)                        // enq
	command = append(command, etx)                        // etx
	command = append(command, calculateChecksum(command)) // BCC
	command = append(command, eot)                        // eot
	command = append(command, 0xFF)                       // eot
	command = append(command, 0xFF)                       // eot

	connection.Write(command)

	sendEnd(connection, address)
}

func sendTextCommand(code int, data string, connection *connection, address int) {
	var command = []byte{soh, (byte)(0x50 + address), 0x70}
	command = append(command, calculateChecksum(command)) // BCA
	command = append(command, stx)                        // stx
	command = addCode(command, code)                      // code
	command = addTextData(command, data)                  // data
	command = append(command, enq)                        // enq
	command = append(command, etx)                        // etx
	command = append(command, calculateChecksum(command)) // BCC
	command = append(command, eot)                        // eot
	command = append(command, 0xFF)                       // eot
	command = append(command, 0xFF)                       // eot

	connection.Write(command)

	sendEnd(connection, address)
}

func atEndOfPacket(reply []byte) bool {
	replyLen := len(reply)
	if replyLen > 3 {
		return reply[replyLen-3] == eot &&
			reply[replyLen-2] == 0xFF &&
			reply[replyLen-1] == 0xFF
	}
	return false
}

func getGetMessageCode(reply []byte) int {
	code := uint(0)
	offset := uint(0)
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == codeOffset {
			value := uint(reply[i] & 0x0F)
			value <<= offset
			code += value
			offset += 4
		}
	}

	return int(code)
}

func getMessageString(reply []byte) string {
	var data string
	var offset uint
	var tempValue byte
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == dataOffset {
			value := reply[i] & 0xF
			value <<= offset
			tempValue += value
			if offset == 4 {
				if tempValue != 0 {
					data += string(tempValue)
				}

				offset = 0
				tempValue = 0
			} else {
				offset = 4
			}
		}
	}

	return data
}

func getMessageDataArray(reply []byte, bits uint) []int {
	var replyItems []int
	var offset uint
	var tempValue int
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == dataOffset {
			var value = reply[i] & 0x0F
			value <<= offset
			tempValue += int(value)
			if offset == bits {
				replyItems = append(replyItems, tempValue)
				tempValue = 0
				offset = 0
			} else {
				offset += 4
			}
		}
	}

	if offset != 0 {
		replyItems = append(replyItems, tempValue)
	}

	return replyItems
}

func getMessageDataShortArray(reply []byte) []int {
	return getMessageDataArray(reply, 12)
}

func getMessageDataTwoByteArray(reply []byte) []int {
	return getMessageDataArray(reply, 4)
}

func getMessageData(reply []byte) int {
	var data int
	var offset uint
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == dataOffset {
			value := reply[i] & 0x0F
			value <<= offset
			data += int(value)
			offset += 4
		}
	}

	return data
}

func getMessageBytes(reply []byte) []byte {
	var replyItems []byte
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == dataOffset {
			replyItems = append(replyItems, reply[i]&0x0F)
		}
	}

	return replyItems
}

func getMessageBools(reply []byte) []bool {
	var replyItems []bool
	for i := 5; i < len(reply); i++ {
		if reply[i] == etx {
			break
		}

		if (reply[i] & 0xF0) == dataOffset {
			replyItems = append(replyItems, (reply[i]&0x1) != 0)
			replyItems = append(replyItems, (reply[i]&0x2) != 0)
			replyItems = append(replyItems, (reply[i]&0x4) != 0)
			replyItems = append(replyItems, (reply[i]&0x8) != 0)
		}
	}

	return replyItems
}
