package server

import (
	e "../server/environment"
	"os"
	"strconv"
)

const (
	portVariable           = "g-server-port"
	defaultPort            = 9000
	environmentVariableKey = "environment-identifier"
)

var env *e.Environment

func getPort() int {
	port := getVariable(portVariable)

	if len(port) > 0 {
		p, _ := strconv.Atoi(port)
		return p
	}

	return defaultPort
}

func createEnvironment(identifier int) {
	port := getPort()

	env = &e.Environment{
		Port:       port,
		PortForGin: ":" + strconv.Itoa(port),
	}

	if identifier != e.Test && identifier != e.Development && identifier != e.Production {
		env.Identifier = e.Development
	} else {
		env.Identifier = identifier
	}
}

func getVariable(variable string) string {
	return os.Getenv(variable)
}