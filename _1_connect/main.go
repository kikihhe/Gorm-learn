package connect

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=iaas_db_user password=iaas_db_pass dbname=iaas_db port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败:", err)
	} else {
		fmt.Println("数据库连接成功:", db)
	}
	Db = db
}

func main() {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=iaas_db_user password=iaas_db_pass dbname=iaas_db port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败:", err)
	} else {
		fmt.Println("数据库连接成功:", db)
	}
	Db = db
}
