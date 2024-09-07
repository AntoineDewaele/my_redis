package cmd_handler

import (
	"reflect"
	"strings"
	"strconv"
)

type Commands struct {}

func HandleCommand(cmd string, args ...string) (string, string) {
	commands := Commands{}

	cmd = toMethodName(cmd)

	method := reflect.ValueOf(commands).MethodByName(cmd)

	if (method.IsValid()) {
		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
		}
		
		// Call the method
		result := method.Call(reflectArgs)
		
		// Assuming the method returns a single string value
		return result[0].String(), result[1].String()
	}

	return "", "unknown command"+cmd
}

// Converts an uppercase command to CamelCase method name
func toMethodName(cmd string) string {
	// Convert the command to title case
	parts := strings.Split(cmd, " ")
	for i, part := range parts {
			if len(part) > 0 {
					parts[i] = strings.Title(strings.ToLower(part))
			}
	}
	return strings.Join(parts, "")
}

// Parses the options of a command
func parseOptions(args []string, optionsAvailable map[string]map[string]interface{}) (map[string][]interface{}, string) {
	options := make(map[string][]interface{})
	for i := 0; i < len(args); i++ {
		arg := args[i]
		option, ok := optionsAvailable[arg]
		if !ok {
			return nil, "syntax error"
		}

		argNb, _ := option["argNb"].(int)
		if argNb == 0 {
			options[arg] = []interface{}{}
			continue
		}

		types, _ := option["types"].([]string)
		optionArgs, err := parseArgs(args, &i, argNb, types)
		if err != "" {
			return nil, err
		}
		options[arg] = optionArgs
	}

	return options, ""
}

func parseArgs(args []string, index *int, argNb int, types []string) ([]interface{}, string) {
	optionArgs := make([]interface{}, 0, argNb)

	for j := 0; j < argNb; j++ {
		(*index)++
		if *index >= len(args) {
			return nil, "syntax error"
		}

		argValue, err := convertArg(args[*index], types[j])
		if err != "" {
			return nil, err
		}
		optionArgs = append(optionArgs, argValue)
	}

	return optionArgs, ""
}

func convertArg(value string, expectedType string) (interface{}, string) {
	switch expectedType {
	case "int":
		argValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, "value is not an integer or out of range"
		}
		return argValue, ""
	default:
		return nil, "unknown type"
	}
}
