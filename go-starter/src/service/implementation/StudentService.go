package implementation

import (
	"../../database"
	"../../database/entity"
)

type StudentService struct {

}

func GetStudents() []entity.Student{
	students := make([]entity.Student, 0)
	_ = database.Engine.Find(&students)
	return students
}
