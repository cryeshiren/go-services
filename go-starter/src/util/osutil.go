package util

import "os"

func GetVariable(variable string) string {
	return os.Getenv(variable)
}
