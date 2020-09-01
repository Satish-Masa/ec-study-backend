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
	defer db.Close()

	jpfaker := domainItem.Item{Name: "吾輩は猫である", Description: "吾輩は猫である。名前はまだ無い。どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙を", Price: 600, Stock: 100}
	db.Create(&jpfaker)
	jpfaker = domainItem.Item{Name: "夜明け前", Description: "木曾路はすべて山の中である。あるところは岨づたいに行く崖の道であり、あるところは数十間の深さに臨む木曾川の岸であり、あるところは山の尾をめぐる谷の入り口である。一筋の街道はこの深い森林地帯を貫いていた。東ざかいの桜沢から、西の十曲峠まで、木曾十一宿はこの街道に添うて、二十二里余にわたる長い谿谷の間に散在していた。道路の位置も幾たびか改まったもので、古道はいつのまにか深い山間に埋もれた。名高い桟も、蔦のかずらを頼みにしたような危い場処ではなくなって、徳川時代の末にはすでに渡ることのできる橋であった。新規に新規にとできた道はだんだん谷の下の方の位置へと降って来た。道の狭いところには、木を伐って並べ、藤づるでからめ、それで街道の狭いのを補った。長い間にこの木曾路に起こって来た変化は、いくらかずつでも嶮岨な山坂の多いところを歩きよくした。そのかわり、大雨ごとにやって来る河水の氾濫が旅行を困難にする", Price: 700, Stock: 70}
	db.Create(&jpfaker)
	jpfaker = domainItem.Item{Name: "人間失格", Description: "恥の多い生涯を送って来ました。自分には、人間の生活というものが、見当つかないのです。自分は東北の田舎に生れましたので、汽車をはじめて見たのは、よほど大きくなってからでした。自分は停車場のブリッジを、上って、降りて、そうしてそれが線路をまたぎ越えるために造られたものだという事には全然気づかず、ただそれは停車場の構内を外国の遊戯場みたいに、複雑に楽しく、ハイカラにするためにのみ、設備せられてあるものだとばかり思っていました。しかも、かなり永い間そう思っていたのです。ブリッジの上ったり降りたりは、自分にはむしろ、ずいぶん垢抜けのした遊戯で、それは鉄道のサーヴィスの中でも、最も気のきいたサーヴィスの一つだと思っていたのですが、のちにそれはただ旅客が線路をまたぎ越えるための頗る実利的な階段に過ぎないのを発見して、にわかに興が覚めました。また、自分は子供の頃、絵本で地下鉄道というものを見て、これもやは", Price: 30, Stock: 40}
	db.Create(&jpfaker)

	f := faker.New()
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		item := domainItem.Item{Name: f.Company().Name(), Description: f.Lorem().Text(255), Price: rand.Intn(100000), Stock: rand.Intn(100)}
		db.Create(&item)
	}

	user := infrastructure.NewUserRepository(db)
	item := infrastructure.NewItemRepository(db)
	cart := infrastructure.NewCartRepository(db)
	ordered := infrastructure.NewOrderedRepository(db)
	mail := infrastructure.NewMailRepository(db)
	rest := &interfaces.Rest{
		UserRepository:  user,
		ItemRepository:  item,
		CartRepository:  cart,
		OrderRepository: ordered,
		MailRepository:  mail,
	}
	rest.Start()
}
