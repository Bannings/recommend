package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"recommend/golbal"
	"sync"
)

var (
	dbUserfeature *gorm.DB
	featureOnce   sync.Once
	dbRecall      *gorm.DB
	recallOnce    sync.Once
	dbBanner      *gorm.DB
	bannerOnce    sync.Once
	dbComicData   *gorm.DB
	comicDataOnce sync.Once
)

func OpenMysql(dbConfig golbal.DBConfig) (*gorm.DB, error) {

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&allowNativePasswords=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema)

	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(30)
	db.DB().SetMaxIdleConns(10)
	return db, nil
}

//func CustomizeDB() *gorm.DB {
//	customizeOnce.Do(func() {
//		conf := gobal.GetConfig()
//		userConf := gobal.DBConfig{Schema: conf.DbCustomize.Schema, Username: conf.DbCustomize.Username, Password: conf.DbCustomize.Password, Host: conf.DbCustomize.Host, Port: conf.DbCustomize.Port}
//		db, err := OpenMysql(userConf)
//		if err != nil {
//			panic(err)
//		}
//		dbCustomize = db
//	})
//
//	return dbCustomize
//}

func UserFeatureDB() *gorm.DB {
	featureOnce.Do(func() {
		conf := golbal.GetConfig()
		userConf := golbal.DBConfig{Schema: conf.DbFeatures.Schema, Username: conf.DbFeatures.Username, Password: conf.DbFeatures.Password, Host: conf.DbFeatures.Host, Port: conf.DbFeatures.Port}
		db, err := OpenMysql(userConf)
		if err != nil {
			panic(err)
		}
		dbUserfeature = db
	})

	return dbUserfeature
}

func RecallDB() *gorm.DB {
	recallOnce.Do(func() {
		conf := golbal.GetConfig()
		userConf := golbal.DBConfig{Schema: conf.DbRecall.Schema, Username: conf.DbRecall.Username, Password: conf.DbRecall.Password, Host: conf.DbRecall.Host, Port: conf.DbRecall.Port}
		db, err := OpenMysql(userConf)
		if err != nil {
			panic(err)
		}
		dbRecall = db
	})

	return dbRecall
}

func BannerDB() *gorm.DB {
	bannerOnce.Do(func() {
		conf := golbal.GetConfig()
		userConf := golbal.DBConfig{Schema: conf.DbBanner.Schema, Username: conf.DbBanner.Username, Password: conf.DbBanner.Password, Host: conf.DbBanner.Host, Port: conf.DbBanner.Port}
		db, err := OpenMysql(userConf)
		if err != nil {
			panic(err)
		}
		dbBanner = db
	})

	return dbBanner
}

func ComicDataDB() *gorm.DB {
	comicDataOnce.Do(func() {
		conf := golbal.GetConfig()
		userConf := golbal.DBConfig{Schema: conf.DBComic.Schema, Username: conf.DBComic.Username, Password: conf.DBComic.Password, Host: conf.DBComic.Host, Port: conf.DBComic.Port}
		db, err := OpenMysql(userConf)
		if err != nil {
			panic(err)
		}
		dbComicData = db
	})

	return dbComicData
}
