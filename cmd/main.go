package main

import (
	"my_redis/internal/server"
	"my_redis/internal/reader"
	"net"
	"fmt"
	"strings"
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
	reader := reader.New(conn)

	for {
		cmd, err := reader.ReadCmd()

		if err != nil {
			return
		}

		fmt.Println(cmd)
			
		if strings.Contains(cmd, "PING") {
			s.Write(conn, "+PONG\r\n")
		} else {
			s.Write(conn, "-ERR unknown command"+cmd+"\r\n")
		}
	}
}