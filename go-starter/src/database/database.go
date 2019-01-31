package database

import (
	"../database/entity"
	"../util"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

const(
	postgreUserVariableKey = "postgre-user"
	postgrePasswordVariableKey = "postgre-password"
	postgreDBNameVariableKey = "postgre-password"
	defaultConnectionString = "user=postgres dbname=performance sslmode=disable"
)

var Engine *xorm.Engine

func InitEngine() {
	e, err := xorm.NewEngine("postgres", getConnecionString())

	if err != nil {
		log.Fatal(err)
	}

	Engine = e
}

func SyncTable(entityPackage string) {
	_ = Engine.Sync2(new(entity.Student))
}

func getConnecionString() string{
	postgreUser := util.GetVariable(postgreUserVariableKey)
	postgrePassword := util.GetVariable(postgrePasswordVariableKey)
	dbName := util.GetVariable(postgreDBNameVariableKey)

	if len(postgreUser) > 0 && len(postgrePassword) > 0 && len(dbName) > 0{
		return "user=" + postgreUser + " dbname=" + dbName + " sslmode=disable"
	}

	return defaultConnectionString
}