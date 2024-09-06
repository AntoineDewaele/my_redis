package main

import (
	"my_redis/internal/server"
	"net"
)

func main() {
	s := server.New()
	s.Start()
	defer s.Stop()

	for {
		go handleConnection(s, s.AcceptConnection())
	}
}

func handleConnection(s *server.Server, conn net.Conn) {
	defer conn.Close()

	for {
		cmd, err := s.Read(conn)
		if err != nil {
			return
		}

		s.Write(conn, cmd)
	}
}