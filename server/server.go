package server

import (
	"fmt"
	"io"

	"github.com/kennardpeters/ExampleGoServer/datastore"
	"golang.org/x/exp/rand"
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

		first := rand.ExpFloat64()
		second := rand.ExpFloat64()

		if first > second {
	
			email, err := s.store.SelectEmailByUserID(ws.Request().Context(), string(msg))
			if err != nil {
				s.broadcast([]byte(err.Error()))
				return
			}
			s.broadcast([]byte(email))
		} else {

			link, err := s.store.SelectLinksByUserID(ws.Request().Context(), string(msg))
			if err != nil {
				s.broadcast([]byte(err.Error()))
				return
			}
			s.broadcast([]byte(link))
		}


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
