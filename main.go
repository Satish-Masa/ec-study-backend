package main

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/ec-backend/config"
	"github.com/Satish-Masa/ec-backend/infrastructure"
	"github.com/Satish-Masa/ec-backend/interfaces"
	"github.com/gchaincl/dotsql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.Init()
}

func main() {
	tmp := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dot, err := dotsql.LoadFromFile("db/create_table.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Exec(db.DB(), "create-users-table")
	if err != nil {
		log.Fatal(err)
	}

	user := infrastructure.NewUserRepository(db)
	rest := &interfaces.Rest{UserRepository: user}
	rest.Start()
}
