package database

import (
	"fmt"
	_"log"
	_"strconv"

	"github.com/kahono922/scanpay/config"
	"github.com/kahono922/scanpay/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect() *gorm.DB{
	var err error
	var DB *gorm.DB
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",config.Config("DB_USER"),config.Config("DB_PASSWORD"),"tcp",config.Config("DB_HOST"),config.Config("DB_PORT"),config.Config("DB_NAME"))

	DB , err = gorm.Open(mysql.Open(dsn))
	if(err!=nil){
		panic("connection failed")
	}

	fmt.Println("Connection Opened")
	DB.AutoMigrate(&model.Students{})
	return(DB)
}
