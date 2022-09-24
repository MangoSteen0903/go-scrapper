package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RepositoryRequest struct {
	Title       string
	Description string
	Badges      []string
	Language    string
	Stars       int
}

var replacer = strings.NewReplacer(
	"\n", "",
	"  ", " ",
	",", "",
)

var repositories []RepositoryRequest
var Repository = RepositoryRequest{}

func GetRepositories(url string, ch2 chan<- []RepositoryRequest) {
	fmt.Printf("Scrapping Url: %s...\n", url)
	doc := GetBody(url)
	doc.Find("li.Box-row").Each(extractText)
	ch2 <- repositories
}

func extractText(i int, s *goquery.Selection) {

	title := CleanString(s.Find("a.d-inline-block").Text())

	description := CleanString(s.Find("p.wb-break-word").Text())

	badges := CleanBadgeString(s.Find("a.topic-tag").Text())

	language := CleanString(s.Find("div.mt-2 span.mr-3").Eq(0).Text())

	starStr := CleanString(s.Find("div.mt-2 a.no-wrap").Eq(0).Text())

	stars, err := strconv.Atoi(replacer.Replace(starStr))
	HandleErr(err)
	Repository = RepositoryRequest{
		Title:       title,
		Description: description,
		Badges:      badges,
		Language:    language,
		Stars:       stars,
	}
	repositories = append(repositories, Repository)
}
