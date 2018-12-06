package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/cjburchell/profilux-go/protocol"
)

type httpProtocol struct {
	address string
}

// NewProtocol creates a http protocol
func NewProtocol(address string, port int) (protocol.IProtocol, error) {
	var p httpProtocol

	p.address = fmt.Sprintf("%s:%d", address, port)

	return &p, nil
}

func (p *httpProtocol) SendData(code, data int) error {
	url := fmt.Sprintf("http://%s/communication.php?dir=sel&code=%d&data=%d", p.address, code, data)

	resp, err := http.Get(url)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if string(reply) == "Access Denied" {
		return errors.WithStack(errors.New("access denied"))
	}

	command, err := getCommandFromReply(reply)
	if err != nil {
		return errors.WithStack(err)
	}

	dataParam := getDataFromReply(reply)

	if command != code {
		return errors.WithStack(fmt.Errorf("unexpected comand reply: %d", command))
	}

	if strings.HasPrefix(dataParam, "NACK") {
		return errors.WithStack(fmt.Errorf("error in command: %d Reply: %s", command, string(reply)))
	}

	if !strings.HasPrefix(dataParam, "ACK") {
		return errors.WithStack(errors.New("unexpected message: Missing ACK"))
	}

	return nil
}

func getCommandFromReply(reply []byte) (int, error) {
	return strconv.Atoi(strings.SplitN(strings.SplitN(string(reply), "&", 2)[0], "=", 2)[1])
}

func getDataFromReply(reply []byte) string {
	return strings.SplitN(strings.SplitN(string(reply), "&", 2)[1], "=", 2)[1]
}

func (p *httpProtocol) GetDataText(code int) (string, error) {
	reply, err := p.getRawData(code)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(*reply), nil
}

func (p *httpProtocol) GetData(code int) (int, error) {
	data, err := p.getRawData(code)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(*data)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return result, nil
}

func (p *httpProtocol) getRawData(code int) (*string, error) {
	url := fmt.Sprintf("http://%s/communication.php?dir=enq&code=%d", p.address, code)

	//log.Debugf("Get %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	command, err := getCommandFromReply(reply)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dataParam := getDataFromReply(reply)

	if command != code {
		return nil, errors.WithStack(fmt.Errorf("unexpected comand reply: %d, %s", command, reply))
	}

	if strings.HasPrefix(dataParam, "NACK") {
		return nil, errors.WithStack(fmt.Errorf("error in command: %d", command))
	}

	return &dataParam, nil
}

func (p *httpProtocol) Disconnect() {
}
