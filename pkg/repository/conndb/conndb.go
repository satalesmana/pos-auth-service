package conndb

import (
	"fmt"

	"go-grpc-auth-svc/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConn(conf models.MySQLConfiguration) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error when connecting:", err.Error())
	}

	return db, nil
}
