package cmd_handler

import (
	"testing"
	"my_redis/internal/store"
)

func TestSetShouldReturnOk(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := commands.Set("toto", "tata")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldReturnErrorIfTooFewArguments(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto")
	except := "ERR wrong number of arguments for command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetWithGetOptionShouldReturnLastValue(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	commands.Set("toto", "tata")
	result, _ := commands.Set("toto", "titi", "GET")
	except := "tata"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetWithGetOptionWithoutLastValueShouldReturnNil(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, err := commands.Set("toto", "tata", "GET")

	if err != "" || result != "" {
		t.Errorf("Error was incorrect, got: %s and %s, want empty for both.", result, err)
	}
}

func TestSetShouldReturnErrorIfOptionNotValid(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "INVALID")
	except := "syntax error"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}