package cmd_handler

import (
	"my_redis/internal/store"
)

func (_ Commands) Get(args ...string) (string, string) {
	store := store.GetStore()
	res := store.Get(args[0])

	if res == "" {
		return "", ""
	}

	return res, ""
}