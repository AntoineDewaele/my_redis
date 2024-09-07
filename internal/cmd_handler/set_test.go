package cmd_handler

import (
	"testing"
	"my_redis/internal/store"
	"time"
	"strconv"
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

func TestSetShouldReturnErrorIfOptionValueNotValid(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "EX", "INVALID")
	except := "value is not an integer or out of range"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNegative(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "EX", "-1")
	except := "invalid expire time in 'set' command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNotAnInteger(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "EX", "tata")
	except := "value is not an integer or out of range"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNotAnIntegerForPX(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "PX", "tata")
	except := "value is not an integer or out of range"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNotAnIntegerForEXAT(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "EXAT", "tata")
	except := "value is not an integer or out of range"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNotAnIntegerForPXAT(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "PXAT", "tata")
	except := "value is not an integer or out of range"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfOptionValueNegativeForPX(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "PX", "-1")
	except := "invalid expire time in 'set' command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldNotSetIfNXAndKeyExists(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	commands.Set("toto", "tata")
	result, _ := commands.Set("toto", "titi", "NX")
	except := ""

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldSetIfNXAndKeyDoesNotExist(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := commands.Set("toto", "titi", "NX")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldNotSetIfXXAndKeyDoesNotExist(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := commands.Set("toto", "titi", "XX")
	except := ""

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldSetIfXXAndKeyExists(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	commands.Set("toto", "tata")
	result, _ := commands.Set("toto", "titi", "XX")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldSetWithTTLAt(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	ttlNow := time.Now().Add(time.Second).Unix()
	result, _ := 	commands.Set("toto", "tata", "EXAT", strconv.FormatInt(ttlNow, 10))
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}

	if store.GetStore().GetTTL("toto") != ttlNow {
		t.Errorf("Result was incorrect, got: %d, want: %d.", store.GetStore().GetTTL("toto"), ttlNow)
	}
}

func TestSetShouldSetWithTTLAtInMilliseconds(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	ttlNow := time.Now().Unix()
	result, _ := 	commands.Set("toto", "tata", "PXAT", strconv.FormatInt(ttlNow * 1000, 10))
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}

	time.Sleep(50 * time.Millisecond)

	if store.GetStore().GetTTL("toto") != ttlNow {
		t.Errorf("Result was incorrect, got: %d, want: %d.", store.GetStore().GetTTL("toto"), ttlNow)
	}
}

func TestSetShouldReturnErrorIfTTLAtIsInThePast(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	ttlNow := time.Now().Add(-time.Second).Unix()
	_, err := commands.Set("toto", "tata", "EXAT", strconv.FormatInt(ttlNow, 10))
	except := "invalid expire time in 'set' command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldReturnErrorIfTTLAtIsInThePastInMilliseconds(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	ttlNow := time.Now().Add(time.Duration(4) * -time.Second).Unix()
	_, err := commands.Set("toto", "tata", "PXAT", strconv.FormatInt(ttlNow * 1000, 10))
	except := "invalid expire time in 'set' command"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}

func TestSetShouldSetWithTTL(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := 	commands.Set("toto", "tata", "EX", "1")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldSetWithTTLInMilliseconds(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := 	commands.Set("toto", "tata", "PX", "1000")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetShouldSetWithTTLInMillisecondsAndExpire(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	result, _ := 	commands.Set("toto", "tata", "EX", "1")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestSetWithKeepTTLShouldKeepTTL(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	ttl := time.Now().Add(time.Duration(4) * time.Second).Unix()
	commands.Set("toto", "tata", "EXAT", strconv.FormatInt(ttl, 10))
	result, _ := 	commands.Set("toto", "titi", "KEEPTTL")
	except := "OK"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}

	if store.GetStore().GetTTL("toto") != ttl {
		t.Errorf("Result was incorrect, got: %d, want: %d.", store.GetStore().GetTTL("toto"), ttl)
	}
}

func TestSetWithNonDuplicateOptionsShouldReturnError(t *testing.T) {
	store.GetStore().Reset()
	commands := Commands{}
	_, err := commands.Set("toto", "tata", "EX", "1", "KEEPTTL")
	except := "syntax error"

	if err != except {
		t.Errorf("Error was incorrect, got: %s, want: %s.", err, except)
	}
}
