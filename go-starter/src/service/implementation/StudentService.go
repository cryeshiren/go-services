package implementation

import (
	"../../database"
	"../../database/entity"
)

type StudentService struct {

}

func GetStudentById(id int) *entity.Student{
	student := new(entity.Student)
	database.Engine.Id(id).Get(student)
	return student
}
