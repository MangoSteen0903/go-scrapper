package utils

import (
	"fmt"
	"strings"

	ccsv "github.com/tsak/concurrent-csv-writer"
)

func WriteCsv(writer *ccsv.CsvWriter, repository RepositoryRequest, ch chan<- bool) {
	writer.Write([]string{
		repository.Title,
		repository.Description,
		repository.Language,
		strings.Join(repository.Badges, ","),
		fmt.Sprintf("%d", repository.Stars)},
	)
	ch <- true
}
