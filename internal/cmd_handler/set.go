package cmd_handler

import (
	"my_redis/internal/store"
	"fmt"
)

func (_ Commands) Set(args ...string) string {
	
	optionsAvailable := map[string]map[string]interface{}{
		"EX": {
			"argNb": 1,
			"types": []string{"int"},
		},
		"PX": {
			"argNb": 1,
			"types": []string{"int"},
		},
		"EXAT": {
			"argNb": 1,
			"types": []string{"int"},
		},
		"PXAT": {
			"argNb": 1,
			"types": []string{"int"},
		},
		"NX": {
			"argNb": 0,
		},
		"XX": {
			"argNb": 0,
		},
		"KEEPTL": {
			"argNb": 0,
		},
		"GET": {
			"argNb": 0,
		},
	}

	options, err := parseOptions(args[2:], optionsAvailable)
	if err != "" {
		return err;
	}

	fmt.Println(options)

	store := store.GetStore()
	store.Set(args[0], args[1])

	return "+OK\r\n"
}