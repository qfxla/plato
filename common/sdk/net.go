package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/hardcore-os/plato/common/tcp"
	"net"
)

type connect struct {
	conn               *net.TCPConn
	sendChan, recvChan chan *Message
}

func newConnet(ip net.IP, port int) *connect {
	clientConn := &connect{
		sendChan: make(chan *Message),
		recvChan: make(chan *Message),
	}
	addr := &net.TCPAddr{IP: ip, Port: port}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("DialTCP.err=%+v", err)
		return nil
	}
	clientConn.conn = conn
	go func() {
		for {
			data, err := tcp.ReadData(conn)
			if err != nil {
				fmt.Printf("ReadData err:%+v", err)
			}
			msg := &Message{}
			json.Unmarshal(data, msg)
			clientConn.recvChan <- msg
		}
	}()
	return clientConn
}

func (c *connect) send(data *Message) {
	// 直接发送给接收方
	c.recvChan <- data
}

func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

func (c *connect) close() {
	// 目前没啥值得回收的
}
