package webSocketClient

import (
	"github.com/gorilla/websocket"
	"github.com/vasvatoskin/CLIgin/internal/shared"
	"log"
)

type Client struct {
	conn         *websocket.Conn
	incomingChan chan shared.ServerMessage
	outgoingChan chan shared.ClientMessage
	id           uint64
	close        chan bool
}

func New() *Client {
	return &Client{
		incomingChan: make(chan shared.ServerMessage),
		outgoingChan: make(chan shared.ClientMessage),
		close:        make(chan bool),
	}
}

func (c *Client) Connect(serverURL string) {
	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Println("Connect create error: ", err)
		return
	}
}

func (c *Client) Disconnect() {
	msg := shared.ClientMessage{Type: shared.DisconnectMessage}
	c.conn.WriteJSON(msg)
	c.conn.Close()
	close(c.close)
}

func (c *Client) SendServerMsg() {
	defer close(c.outgoingChan)

	var msg shared.ClientMessage

	for {
		select {
		case <-c.close:
			log.Println("Sender closed")
			return
		case msg = <-c.outgoingChan:
			msg.ID = c.id
			err := c.conn.WriteJSON(msg)
			if err != nil {
				log.Println("Error while sending message: ", err)
			}
		}
	}

}

func (c *Client) ReceiveServerMsg() {
	defer close(c.incomingChan)

	var msg shared.ServerMessage

	for {
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Receiver closed")
			return
		}
		if msg.Type == shared.WelcomeMessage {
			c.id = msg.ID
		}
		c.incomingChan <- msg
	}
}

func (c *Client) GetOutgoingChannel() chan shared.ClientMessage {
	return c.outgoingChan
}

func (c *Client) GetIncomingChannel() chan shared.ServerMessage {
	return c.incomingChan
}
