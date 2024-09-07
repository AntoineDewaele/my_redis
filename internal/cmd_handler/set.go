package cmd_handler

import (
	"my_redis/internal/store"
	"time"
)

func (c Commands) Set(args ...string) (string, string) {
	
	optionsAvailable := c.getOptionsAvailable()

	if len(args) < 2 {
		return "", "ERR wrong number of arguments for command"
	}

	options, err := parseOptions(args[2:], optionsAvailable)
	if err != "" {
		return "", err;
	}

	if err := c.checkOptionsUnicity(options); err != "" {
		return "", err
	}

	returnValue, ttl, canSet, err := c.handleOptions(options, args);
	if err != "" {
		return "", err
	}
	
	if canSet {
		store.GetStore().SetWithTTL(args[0], args[1], ttl)
	}

	return returnValue, err
}

func (_ Commands) handleOptions(options map[string][]interface{}, args []string) (string, int64, bool, string) {
	store := store.GetStore()
	var ttl int64
	returnValue := "OK"
	canSet := true

	for option, _ := range options {
		
		switch option {
	
			case "EX":
				if options["EX"][0].(int) < 0 {
					return "", 0, false, "invalid expire time in 'set' command"
				}
				ttl = time.Now().Add(time.Duration(int64(options["EX"][0].(int))) * time.Second).Unix()
		
			case "PX":
				if options["PX"][0].(int) < 0 {
					return "", 0, false, "invalid expire time in 'set' command"
				}
				ttl = time.Now().Add(time.Duration(int64(options["PX"][0].(int))) * time.Millisecond).Unix()
			
			case "EXAT":
				if time.Now().Unix() > int64(options["EXAT"][0].(int)) {
					return "", 0, false, "invalid expire time in 'set' command"
				}
				ttl = int64(options["EXAT"][0].(int))
			
			case "PXAT":
				if time.Now().Unix() > int64(options["PXAT"][0].(int) / 1000) {
					return "", 0, false, "invalid expire time in 'set' command"
				}
				ttl = int64(options["PXAT"][0].(int)) / 1000
			
			case "NX":
				if store.Get(args[0]) != "" {
					canSet = false
					returnValue = ""
				}
			
			case "XX":
				if store.Get(args[0]) == "" {
					canSet = false
					returnValue = ""
				}
			
			case "KEEPTTL":
				ttl = store.GetTTL(args[0])
			
			case "GET":
				returnValue = store.Get(args[0])
		}
	}

	return returnValue, ttl, canSet,""
}

func (_ Commands) checkOptionsUnicity(options map[string][]interface{}) string {
	nonDuplicateOptions := []string{"EX", "PX", "EXAT", "PXAT", "KEEPTTL"}
	count := 0

	for _, option := range nonDuplicateOptions {
		if _, exists := options[option]; exists {
			count++
		}
	}

	if count > 1 {
		return "syntax error"
	}

	return ""
}

func (_ Commands) getOptionsAvailable() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
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
		"KEEPTTL": {
			"argNb": 0,
		},
		"GET": {
			"argNb": 0,
		},
	}
}