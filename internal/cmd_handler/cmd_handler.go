package cmd_handler

import (
	"reflect"
	"strings"
)

type Commands struct {}

func HandleCommand(cmd string, args ...string) string {
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
		return result[0].String()
	}

	return "-unknown command"+cmd+"\r\n"
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