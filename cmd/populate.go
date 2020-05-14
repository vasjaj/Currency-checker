package cmd

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/cobra"
	"github.com/vasjaj/Currency-checker/db"
)

func init() {
	rootCmd.AddCommand(populateCmd)
}

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populates DB with latest currency information",
	Long:  "Populates DB with latest currency information via https://www.bank.lv/vk/ecb_rss.xml",
	Run:   populate,
}

func populate(cmd *cobra.Command, args []string) {
	res, err := http.Get("https://www.bank.lv/vk/ecb_rss.xml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var resStruct db.XMLCurrencyResponse
	if err := xml.Unmarshal(data, &resStruct); err != nil {
		panic(err)
	}

	parsedRes := parseResponse(resStruct)

	_ = saveResponseData(parsedRes)
	// fmt.Println(resStruct)
}

func parseResponse(res db.XMLCurrencyResponse) []db.CurrencyInformation {
	var result []db.CurrencyInformation

	for _, item := range res.Channel.Items {
		dt, err := time.Parse(time.RFC1123Z, item.Date)
		if err != nil {
			continue
		}
		ci := strings.Split(item.Description, " ")

		ciLen := len(ci) - 1
		for i := 0; i < ciLen; i = i + 2 {
			result = append(result, db.CurrencyInformation{Name: ci[i], Value: ci[i+1], Date: dt})
		}
	}

	return result
}

func saveResponseData(resData []db.CurrencyInformation) error {
	db, err := db.Connect()
	defer db.Close()
	if err != nil {
		return err
	}

	for _, ci := range resData {
		// could be better in one transaction, but seems like overkill in this case
		db.Create(&ci)
	}

	return nil
}
