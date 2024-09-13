package main

import (
	connect "Gorm-learn/_1_connect"
	"Gorm-learn/model"
	"fmt"
)

func main() {
	//student := model.Student{Id: 1, UserName: "xiaohong", Age: 10}
	//rows := UpdateOrSave(&student)
	//fmt.Println(rows)

	//result := UpdateUserNameById(1, "xiaoming")
	//fmt.Println(result)

	student := model.Student{Id: 1, UserName: "xiaogang", Age: 200}
	result := UpdateById(&student)
	fmt.Println(result)
}

func UpdateOrSave(studnet *model.Student) int {
	result := connect.Db.Save(studnet)
	return int(result.RowsAffected)
}

func UpdateUserNameById(id int, username string) int {
	result := connect.Db.Debug().Model(&model.Student{}).Where("id = ?", id).Update("user_name", username)
	return int(result.RowsAffected)
}

func UpdateAgeById(student *model.Student) int {
	result := connect.Db.Model(&model.Student{}).Select("age").Where("id = ?", student.Id).Updates(student)
	return int(result.RowsAffected)
}
func UpdateByIdOmitAge(student *model.Student) int {
	result := connect.Db.Model(&model.Student{}).Omit("age").Where("id = ?", student.Id).Updates(student)
	return int(result.RowsAffected)
}
func UpdateById(student *model.Student) int {
	result := connect.Db.Model(&model.Student{}).Where("id = ?", student.Id).Updates(student)
	return int(result.RowsAffected)
}
