package server

import (
	"net"
	"fmt"
	"os"
)

type Server struct {
	listener net.Listener
}

func New() *Server {
	return &Server{}
}

func (server *Server) Start() {
	var err error
	server.listener, err = net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind port 6379: ", err.Error())
		os.Exit(1)
	}
}

func (server *Server) AcceptConnection() net.Conn {
	conn, err := server.listener.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection: ", err.Error())
		os.Exit(1)
	}

	return conn
}

func (server *Server) Write(conn net.Conn, resp string) {
	_, err := conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("Failed to write response: ", err.Error())
	}
}

func (server *Server) Stop() {
	server.listener.Close()
}

func (server *Server) CloseConnection(conn net.Conn) {
	conn.Close()
}
