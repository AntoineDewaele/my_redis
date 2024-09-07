package main

import (
	"my_redis/internal/server"
	"my_redis/internal/reader"
	"my_redis/internal/cmd_handler"
	"net"
	"fmt"
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
		cmd, args, err := reader.ReadCmd()

		if err != nil {
			return
		}

		fmt.Println(cmd + " " + fmt.Sprint(args))
		resp, errResp := cmd_handler.HandleCommand(cmd, args...)
		writeMessage := ""

		if (resp != "") {
			writeMessage = "+\""+resp+"\"\r\n"
		} else if (errResp != "") {
			writeMessage = "-"+errResp+"\r\n"
		} else {
			writeMessage = "$-1\r\n"
		}

		s.Write(conn, writeMessage)
	}
}