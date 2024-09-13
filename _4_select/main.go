package main

import (
	connect "Gorm-learn/_1_connect"
	"Gorm-learn/model"
	"fmt"
)

func main() {
	//studentList := selectByUserName("xiaoming")
	//
	//for _, student := range studentList {
	//	fmt.Println(student)
	//}

	//r := selectTotalAgeByUserName("xiaoming")
	//fmt.Println(r)

	total := countByUserName("xiaoming")
	fmt.Println(total)
}

func selectById(id int) *model.Student {
	var student model.Student
	result := connect.Db.First(&student, id)
	if result.Error != nil {
		return nil
	}
	return &student
}
func selectByUserName(username string) []model.Student {
	var students []model.Student
	connect.Db.Where("user_name = ?", username).Find(&students)
	return students
}

func selectByUserNameOrderByFields(username string, fields []string) []model.Student {
	var students []model.Student
	connect.Db.Where("user_name = ?", username).Order(fields).Find(&students)
	return students
}

func selectTotalAgeByUserName(username string) interface{} {
	type result struct {
		UserName string
		TotalAge int
	}
	var r result
	//connect.Db.Table("student").Select("user_name, sum(age) as total_age").Where("user_name = ?", username).Group("user_name").Find(&r)
	connect.Db.Select("user_name, sum(age) as total_age").Where("user_name = ?", username).Group("user_name").Find(&r)
	return r
}

func countByUserName(username string) int {
	var total int64
	connect.Db.Debug().Model(&model.Student{}).Where("user_name = ?", username).Count(&total)
	return int(total)
}
