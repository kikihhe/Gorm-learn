package main

import (
	connect "Gorm-learn/_1_connect"
	"Gorm-learn/model"
	"fmt"
)

func main() {

	err := connect.Db.AutoMigrate(model.Student{})
	if err != nil {
		fmt.Println("表创建失败")
	} else {
		fmt.Println("表创建成功")
	}
}
