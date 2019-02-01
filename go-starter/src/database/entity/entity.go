package entity

import "time"

type Student struct {
	Id int64
	Name string
	CreatedAt time.Time `xorm:"created"`
}


type Class struct {
	Id int64
	Name string
	CreatedAt time.Time `xorm:"created"`
}