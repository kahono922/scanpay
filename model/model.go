package model

import "gorm.io/gorm"

type Students struct{
  gorm.Model
  Id uint64 `gorm:"primaryKey;autoIncrement:true"`
  Reg string
  First_name string
  Last_name string
  Password string
}

type Wallet struct{}
