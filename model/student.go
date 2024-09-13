package model

type Student struct {
	Id       int
	UserName string
	Age      int
}

func (student *Student) TableName() string {
	return "student"
}
