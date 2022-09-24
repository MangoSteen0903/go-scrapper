package main

import (
	"github.com/MangoSteen0903/go-scrapper/utils"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

const baseURL = "https://github.com/orgs/vercel/repositories"

func main() {

	result := utils.ScrapURL(baseURL)
	done := make(chan bool)
	file, err := ccsv.NewCsvWriter("./output.csv")
	utils.HandleErr(err)

	defer file.Close()

	file.Write([]string{
		"Title", "Description", "Language", "Badges", "Stars",
	})
	for _, item := range result {
		go utils.WriteCsv(file, item, done)
		<-done
	}
}
