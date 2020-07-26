package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	domainItem "github.com/Satish-Masa/ec-backend/domain/item"

	"github.com/Satish-Masa/ec-backend/config"
	"github.com/Satish-Masa/ec-backend/infrastructure"
	"github.com/Satish-Masa/ec-backend/interfaces"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	migrate "github.com/rubenv/sql-migrate"
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

	migrations := &migrate.FileMigrationSource{
		Dir: "db",
	}
	_, err = migrate.Exec(db.DB(), driver, migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	defer migrate.Exec(db.DB(), driver, migrations, migrate.Down)
	defer db.Close()

	f := faker.New()
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		item := domainItem.Item{Name: f.Company().Name(), Description: f.Lorem().Text(255), Price: rand.Intn(100000)}
		db.Create(&item)
	}

	user := infrastructure.NewUserRepository(db)
	item := infrastructure.NewItemRepository(db)
	rest := &interfaces.Rest{
		UserRepository: user,
		ItemRepository: item,
	}
	rest.Start()
}
