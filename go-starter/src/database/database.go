package database

import (
	e "../server/environment"
	"../util"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/modern-go/reflect2"
	"log"
)

var DbConfig DBConfig

const(
	postgreUserVariableKey = "postgre-user"
	postgrePasswordVariableKey = "postgre-password"
	postgreDBNameVariableKey = "postgre-password"
)

type DBConfig struct {
	User string
	DBName string
}

type EntityName string

var Engine *xorm.Engine
var EntityContainer = make([]EntityName, 0)

func InitEngine() {
	e, err := xorm.NewEngine("postgres", getConnecionString())

	if err != nil {
		log.Fatal(err)
	}

	Engine = e
}

func SyncTable() {
	var types = make([]interface{}, 0)
	for _, entity := range EntityContainer {
		types = append(types, reflect2.TypeByName(string(entity)))
	}
	err := Engine.Sync2(types)

	if err != nil {
		log.Fatal("Sync table failed.")
	}
}

func InitDBConfig(){
	postgreUser := util.GetVariable(postgreUserVariableKey)
	postgrePassword := util.GetVariable(postgrePasswordVariableKey)
	dbName := util.GetVariable(postgreDBNameVariableKey)

	if len(postgreUser) > 0 && len(postgrePassword) > 0 && len(dbName) > 0{
		DbConfig.User = postgreUser
		DbConfig.DBName  = postgrePassword
	}else{
		DbConfig.User = e.ServerConfig.Section("database").Key("user").String()
		DbConfig.DBName = e.ServerConfig.Section("databse").Key("db").String()
	}
}

func getConnecionString() string{
	return fmt.Sprint("user=s% dbname=s% sslmode=disable", DbConfig.User, DbConfig.DBName)
}
