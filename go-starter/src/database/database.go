package database

import (
	"../database/entity"
	e "../server/environment"
	"../util"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
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

func InitDBEngine() {
	e, err := xorm.NewEngine("postgres", getConnecionString())

	if err != nil {
		log.Fatal(err)
	}

	Engine = e
}

func SyncTable() {
	err := Engine.Sync2(new(entity.Student), new(entity.Class))

	if err != nil {
		log.Println(err.Error())
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
		DbConfig.DBName = e.ServerConfig.Section("database").Key("db").String()
	}
}

func getConnecionString() string{
	return fmt.Sprintf("user=%s dbname=%s sslmode=disable", DbConfig.User, DbConfig.DBName)
}
