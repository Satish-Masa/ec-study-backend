package main

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/ec-backend/config"
	domainUser "github.com/Satish-Masa/ec-backend/domain/user"
	"github.com/Satish-Masa/ec-backend/infrastructure"
	"github.com/Satish-Masa/ec-backend/interfaces"
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

	db.AutoMigrate(&domainUser.User{})

	user := infrastructure.NewUserRepository(db)
	rest := &interfaces.Rest{UserRepository: user}
	rest.Start()
}
