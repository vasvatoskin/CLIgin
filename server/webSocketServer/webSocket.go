package webSocketServer

import (
	"github.com/gorilla/websocket"
	"github.com/vasvatoskin/CLIgin/shared"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	upgrader     websocket.Upgrader
	incomingChan chan shared.ClientMessage
	outgoingChan chan shared.ServerMessage
	connections  map[uint64]*websocket.Conn
	nextID       uint64
	mu           sync.Mutex
}

func NewServer() *Server {
	return &Server{
		upgrader:     websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024},
		incomingChan: make(chan shared.ClientMessage),
		outgoingChan: make(chan shared.ServerMessage),
		connections:  make(map[uint64]*websocket.Conn),
		nextID:       1,
	}
}

func (s *Server) HandleWebSockets(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Connection creation errors: ", err)
		return
	}
	defer conn.Close()

	s.mu.Lock()
	clientID := s.nextID
	s.nextID++
	s.connections[clientID] = conn
	s.mu.Unlock()

	welcomeMsg := shared.ServerMessage{
		Type: shared.WelcomeMessage,
		ID:   clientID,
	}
	err = conn.WriteJSON(welcomeMsg)
	if err != nil {
		log.Println("Error while sending welcome message: ", err)
		s.mu.Lock()
		delete(s.connections, clientID)
		s.mu.Unlock()
		return
	}

	var msg shared.ClientMessage

	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error while receiving message: ", err)
			s.mu.Lock()
			delete(s.connections, clientID)
			s.mu.Unlock()
			return
		}
		if msg.Type == shared.DisconnectMessage {
			log.Printf("Client ID: %d disconected", clientID)
			s.mu.Lock()
			delete(s.connections, clientID)
			s.mu.Unlock()
			return
		}
		s.incomingChan <- msg
	}
}

func (s *Server) HandleBroadcast() {
	var msg shared.ServerMessage

	for {
		msg = <-s.outgoingChan
		for id, conn := range s.connections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("Error while sending message: ", err)
				s.mu.Lock()
				delete(s.connections, id)
				s.mu.Unlock()
			}
		}
	}
}

func (s *Server) Shutdown() {
	msg := shared.ServerMessage{Type: shared.DisconnectMessage}
	s.mu.Lock()
	for id, conn := range s.connections {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error while sending shutdown message to client:", err)
		}
		conn.Close()
		delete(s.connections, id)
	}
	s.mu.Unlock()

	close(s.incomingChan)
	close(s.outgoingChan)
}

func (s *Server) GetOutgoingChannel() chan shared.ServerMessage { return s.outgoingChan }

func (s *Server) GetIncomingChannel() chan shared.ClientMessage { return s.incomingChan }
