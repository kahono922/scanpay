package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kahono922/scanpay/config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


var DB *gorm.DB

func connect(){
	var err error
	p := config.Config("DB_PORT")
	port ,err := strconv.ParseUint(p,10,32)
	
	if(err!=nil){
		log.Println("Invalid DB_PORT")
	}
	

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB , err := gorm.Open(mysql.Open(dsn))
	if(err!=nil){
		panic("connection failed")
	}

	fmt.Println("Connection Opened")
	fmt.Println(DB)
}
