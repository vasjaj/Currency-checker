package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// didn't want to implement full router with persistent db connection, because seems like overkill for this task
// will open and close db connection on each function

func Connect() (*gorm.DB, error) {
	cURL := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", "some_user", "some_password", "currency_db")

	db, err := gorm.Open("mysql", cURL)
	if err != nil {
		return db, err
	}

	db.AutoMigrate(&CurrencyInformation{})

	return db, nil
}

func GetLatestHistory() ([]CurrencyInformation, error) {
	db, err := Connect()
	if err != nil {
		return []CurrencyInformation{}, err
	}
	defer db.Close()

	var latestInfo CurrencyInformation
	if err := db.Order("date desc").First(&latestInfo).Error; err != nil {
		return []CurrencyInformation{}, err
	}

	var infos []CurrencyInformation
	if err := db.Where("date = ?", latestInfo.Date).Find(&infos).Error; err != nil {
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
