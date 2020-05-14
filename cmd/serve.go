package cmd

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/vasjaj/Currency-checker/db"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs service on :8080 port",
	Long:  "",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	e := echo.New()

	e.GET("/", showLatest)
	e.GET("/:currency", showSpecific)

	e.Logger.Fatal(e.Start(":8080"))
}

func showLatest(c echo.Context) error {
	data, err := db.GetLatestHistory()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error ")
	}

	var resData []db.CurrencyInformationResponse

	for _, d := range data {
		resData = append(resData, d.Response())
	}

	return c.JSON(200, resData)
}

func showSpecific(c echo.Context) error {
	data, err := db.GetHistoryByCurrency(c.Param("currency"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error ")
	}

	var resData []db.CurrencyInformationResponse

	for _, d := range data {
		resData = append(resData, d.Response())
	}

	return c.JSON(200, resData)
}
