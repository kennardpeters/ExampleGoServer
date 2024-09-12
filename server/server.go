package server

import (
	"fmt"
	"io"

	"github.com/kennardpeters/ExampleGoServer/datastore"
	"golang.org/x/net/websocket"
)


type Server struct {
	conns map[*websocket.Conn]bool
	store *datastore.DataStore
}

func NewServer(store *datastore.DataStore) *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
		store: store,
	}
}


func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client: ", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}


func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
	
		email, err := s.store.SelectEmailByUserID(string(msg))
		if err != nil {
			s.broadcast([]byte(err.Error()))
			return
		}

		s.broadcast([]byte(email))
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
			fmt.Println("write error:", err)
			}

		}(ws)
	}
}
