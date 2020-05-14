package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type XMLCurrencyResponse struct {
	Channel struct {
		Title       string            `xml:"title"`
		Description string            `xml:"description"`
		Link        string            `xml:"link"`
		Items       []XMLCurrencyItem `xml:"item"`
	} `xml:"channel"`
}

type XMLCurrencyItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Date        string `xml:"pubDate"`
}

type CurrencyInformation struct {
	gorm.Model

	Name  string    `gorm:"not null;unique_index:name_date_index"`
	Value string    `gorm:"not null"`
	Date  time.Time `gorm:"not null;unique_index:name_date_index"`
}

type CurrencyInformationResponse struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func (ci *CurrencyInformation) Response() CurrencyInformationResponse {
	return CurrencyInformationResponse{
		Name:      ci.Name,
		Value:     ci.Value,
		Date:      ci.Date,
		CreatedAt: ci.CreatedAt,
	}
}
