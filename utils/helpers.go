package utils

import "strings"

var service_name string  = "main.go"

func ServiceName(name string) string {
	if name != "" {return strings.ToUpper(name)}
	return service_name
}