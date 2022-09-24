package scrap

import (
	"fmt"

	"github.com/MangoSteen0903/go-scrapper/utils"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

func GitScrapper(company string) {
	gitURL := fmt.Sprintf("https://github.com/orgs/%s/repositories", company)
	result := utils.ScrapURL(gitURL)
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
