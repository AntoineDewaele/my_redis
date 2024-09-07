package cmd_handler

import (
	"my_redis/internal/store"
	"fmt"
)

func (_ Commands) Set(args ...string) (string, string) {
	
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

	if len(args) < 2 {
		return "", "ERR wrong number of arguments for command"
	}

	options, err := parseOptions(args[2:], optionsAvailable)
	if err != "" {
		return "", err;
	}

	returnValue := "OK"
	store := store.GetStore()

	for option, _ := range options {
		switch option {
			case "EX":
				fmt.Println("EX", args)
			case "PX":
				fmt.Println("PX", args)
			case "EXAT":
				fmt.Println("EXAT", args)
			case "PXAT":
				fmt.Println("PXAT", args)
			case "NX":
				fmt.Println("NX", args)
			case "XX":
				fmt.Println("XX", args)
			case "KEEPTL":
				fmt.Println("KEEPTL", args)
			case "GET":
				returnValue = store.Get(args[0])
		}
	}
	
	store.Set(args[0], args[1])

	return returnValue, err
}