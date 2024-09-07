package store

import (
	"testing"
	"time"
)

func TestStoreGetShouldReturnData(t *testing.T) {
	store := GetStore()
	store.Reset()
	store.Set("toto", "tata")
	result := store.Get("toto")
	except := "tata"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestStoreGetShouldReturnEmptyIfKeyDoesNotExist(t *testing.T) {
	store := GetStore()
	store.Reset()
	result := store.Get("toto")

	if result != "" {
		t.Errorf("Result was incorrect, got: %s, want empty.", result)
	}
}

func TestStoreGetTTLShouldReturnTTL(t *testing.T) {
	store := GetStore()
	store.Reset()
	ttl := time.Now().Add(time.Hour).Unix()
	store.SetWithTTL("toto", "tata", ttl)
	result := store.GetTTL("toto")

	if result != ttl {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, ttl)
	}
}

func TestStoreGetTTLShouldReturnZeroIfKeyDoesNotExist(t *testing.T) {
	store := GetStore()
	store.Reset()
	result := store.GetTTL("toto")

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want 0.", result)
	}
}

func TestStoreGetTTLShouldReturnZeroIfKeyHasNoTTL(t *testing.T) {
	store := GetStore()
	store.Reset()
	store.Set("toto", "tata")
	result := store.GetTTL("toto")

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want 0.", result)
	}
}

func TestStoreSetShouldSetData(t *testing.T) {
	store := GetStore()
	store.Reset()
	store.Set("toto", "tata")
	result := store.Get("toto")
	except := "tata"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}
}

func TestStoreSetWithTTLShouldSetDataAndTTL(t *testing.T) {
	store := GetStore()
	store.Reset()
	ttl := time.Now().Add(time.Hour).Unix()
	store.SetWithTTL("toto", "tata", ttl)
	result := store.Get("toto")
	except := "tata"

	if result != except {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, except)
	}

	resultTTL := store.GetTTL("toto")

	if resultTTL != ttl {
		t.Errorf("Result was incorrect, got: %d, want: %d.", resultTTL, ttl)
	}
}

func TestStoreSetWithTTLShouldSetDataAndTTLAndDeleteIfExpired(t *testing.T) {
	store := GetStore()
	store.Reset()
	ttl := time.Now().Add(-time.Hour).Unix()
	store.SetWithTTL("toto", "tata", ttl)
	result := store.Get("toto")

	if result != "" {
		t.Errorf("Result was incorrect, got: %s, want empty.", result)
	}
}

func TestStoreResetShouldDeleteAllData(t *testing.T) {
	store := GetStore()
	store.Reset()
	store.Set("toto", "tata")
	store.Reset()
	result := store.Get("toto")

	if result != "" {
		t.Errorf("Result was incorrect, got: %s, want empty.", result)
	}
}