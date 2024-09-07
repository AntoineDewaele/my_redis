package cmd_handler

import (
	"testing"
)

func TestPingShouldReturnPong(t *testing.T) {
	commands := Commands{}
	result, _ := commands.Ping()
	except := "PONG"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestPingShouldReturnErrorIfTooManyArguments(t *testing.T) {
	commands := Commands{}
	_, err := commands.Ping("toto", "tata")
	except := "wrong number of arguments for 'ping' command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestPingShouldReturnArgumentIfOne(t *testing.T) {
	commands := Commands{}
	result, _ := commands.Ping("toto")
	except := "toto"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}