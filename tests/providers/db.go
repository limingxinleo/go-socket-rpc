package providers

import "github.com/jinzhu/gorm"

type Mysql struct {
	DB *gorm.DB
}

func (mysql *Mysql) Register() {
	db, _ := gorm.Open("mysql", "root@/phalcon?charset=utf8&parseTime=True&loc=Local")

	mysql.DB = db
}

func (mysql *Mysql) GetInstance() *gorm.DB {
	return mysql.DB
}
