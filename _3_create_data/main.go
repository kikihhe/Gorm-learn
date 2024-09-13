package main

import (
	connect "Gorm-learn/_1_connect"
	"Gorm-learn/model"
	"fmt"
)

func main() {

	students := []model.Student{
		{UserName: "xiaowang", Age: 17},
		{UserName: "xiaohong", Age: 22},
		{UserName: "xiaohe", Age: 15},
	}
	rows := insertBatch(&students)
	if rows == len(students) {
		fmt.Println("全部插入成功")
	}
}

func insert(student *model.Student) int {
	result := connect.Db.Create(student)
	return int(result.RowsAffected)
}

func insertBatch(students *[]model.Student) int {
	result := connect.Db.Create(students)
	return int(result.RowsAffected)
}
func insertBatchSize(students *[]model.Student, size int) int {
	result := connect.Db.CreateInBatches(students, size)
	return int(result.RowsAffected)
}

func insertWithFields(student *model.Student, fields ...string) int {
	result := connect.Db.Select(fields).Create(student)
	return int(result.RowsAffected)
}
