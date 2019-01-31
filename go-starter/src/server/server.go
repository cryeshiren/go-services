package server

import (
	r "../controller/router"
	"../database"
	e "../server/environment"
	"../util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var environment e.Environment

func Start() {
	initEnvironment()
	injectionDifferentGinConfigWithEnv(environment.Identifier)
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(getLoggerFormat()))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	//Register router
	r.InitializeRouter(router)
	//init database
	database.InitEngine()
	//start server
	err := router.Run(environment.PortForGin)
	log.Println(err)
}

func initEnvironment() {
	environmentIdentifier := util.GetVariable(e.IdentifierVariableKey)
	var config *ini.File
	configFileName := "config.ini"
	if len(environmentIdentifier) > 0 {
		configFileName = "config." + environmentIdentifier + ".ini"
	}
	config, _ = ini.Load(configFileName)

	port, _ := config.Section("http").Key("port").Int()

	identifier := 0
	switch environmentIdentifier {
	case e.DevelopmentIdentifierName:
		identifier = e.Development
	case e.TestIdentifierName:
		identifier = e.Test
	case e.ProductionIdentifierName:
		identifier = e.Production
	}

	environment = e.Environment{
		PortForGin: ":" + strconv.Itoa(port),
		Port: port,
		Identifier: identifier,
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
