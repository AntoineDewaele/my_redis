package cmd_handler

import (
)

func (_ Commands) Ping(args ...string) (string, string){

	if len(args) > 1 {
		return "", "wrong number of arguments for 'ping' command"
	}

	if len(args) == 1 {
		return args[0], ""
	}

	return "PONG", ""
}