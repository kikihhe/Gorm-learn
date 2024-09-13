package main

import (
	connect "Gorm-learn/_1_connect"
	"Gorm-learn/model"
)

func main() {
	DeleteById(1)
	student := model.Student{Id: 1, UserName: "xiaogang", Age: 200}
	connect.Db.Delete(&student)
	DeleteByUserName("xiaoming")
}

// 删除指定ID的学生记录
func DeleteById(id int) int {
	var student model.Student
	result := connect.Db.Debug().Where("id = ?", id).Delete(&student)
	return int(result.RowsAffected)
}

// 根据用户名删除学生记录
func DeleteByUserName(username string) int {
	var student model.Student
	result := connect.Db.Debug().Where("user_name = ?", username).Delete(&student)
	return int(result.RowsAffected)
}
