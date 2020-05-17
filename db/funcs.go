package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// didn't want to implement full router with persistent db connection, because seems like overkill for this task
// will open and close db connection on each function/request
// could just open db in main and path reference downwards

func Connect() (*gorm.DB, error) {
	cURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	db, err := gorm.Open("mysql", cURL)
	if err != nil {
		return db, err
	}

	db.LogMode(true)

	db.AutoMigrate(&CurrencyInformation{})

	return db, nil
}

func GetLatestHistory() ([]CurrencyInformation, error) {
	db, err := Connect()
	if err != nil {
		return []CurrencyInformation{}, err
	}
	defer db.Close()

	var infos []CurrencyInformation
	if err := db.Where("date >= ?", db.Table("currency_informations").Select("MAX(date)").Where("currency_informations.date IS NOT NULL").SubQuery()).Find(&infos).Error; err != nil {
		return []CurrencyInformation{}, err
	}

	return infos, nil
}

func GetHistoryByCurrency(name string) ([]CurrencyInformation, error) {
	db, err := Connect()
	if err != nil {
		return []CurrencyInformation{}, err
	}
	defer db.Close()

	var infos []CurrencyInformation
	if err := db.Order("date desc").Where("name = ?", name).Find(&infos).Error; err != nil {
		return []CurrencyInformation{}, err
	}

	return infos, nil
}
