package database

import (
	"fmt"
	"github.com/ktruedat/taleBack/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnect() *gorm.DB {
	dsn := "root:admin@tcp(localhost:3306)/taleydb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connected MySQL")
	}
	if err = db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&model.Tale{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&model.Category{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.Region{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.Country{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.Weather{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.TimeOfTheDay{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.References{}); err != nil {
		log.Fatal(err)
	}
	seedCategories(db)
	seedCountry(db)
	seedRegion(db)
	seedWeather(db)
	seedTimeOfDay(db)
	return db
}

func seedCategories(db *gorm.DB) {
	var count int64
	db.Model(&model.Category{}).Count(&count)
	if count == 0 {
		for _, category := range model.CategorySeedData {
			db.Create(&category)
		}
	}
}

func seedCountry(db *gorm.DB) {
	var count int64
	db.Model(&model.Country{}).Count(&count)
	if count == 0 {
		for _, country := range model.CountrySeedData {
			db.Create(&country)
		}
	}
}

func seedRegion(db *gorm.DB) {
	var count int64
	db.Model(&model.Region{}).Count(&count)
	if count == 0 {
		for _, region := range model.RegionSeedData {
			db.Create(&region)
		}
	}
}

func seedTimeOfDay(db *gorm.DB) {
	var count int64
	db.Model(&model.TimeOfTheDay{}).Count(&count)
	if count == 0 {
		for _, time := range model.TimeOfDaySeedData {
			db.Create(&time)
		}
	}
}

func seedWeather(db *gorm.DB) {
	var count int64
	db.Model(&model.Weather{}).Count(&count)
	if count == 0 {
		for _, time := range model.WeatherSeedData {
			db.Create(&time)
		}
	}
}
