package utils

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetBody(url string) *goquery.Document {
	res, err := http.Get(url)
	HandleErr(err)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	HandleErr(err)

	return doc
}
