package server

import (
	e "../server/environment"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func Start() {
	initEnvironment()
	injectionDifferentGinConfigWithEnv(env.Identifier)
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(getLoggerFormat()))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	err := router.Run(env.PortForGin)
	log.Println(err)
}

func initEnvironment() {
	environmentIdentifier := getVariable(environmentVariableKey)

	if len(environmentIdentifier) == 0 || strings.EqualFold(environmentIdentifier, e.DevelopmentIdentifierName) {
		createEnvironment(e.Development)
	} else {
		createEnvironment(e.Production)
	}
}

func injectionDifferentGinConfigWithEnv(identifier int) {
	if identifier == e.Development {
		//config sth
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
		// Use the following code if you need to write the logs to file and console at the same time.
		// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}

func getLoggerFormat() func(param gin.LogFormatterParams) string {
	return func(param gin.LogFormatterParams) string {
		//custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}
