package utils

import "strconv"

func GetPage(url string) int {
	var totalPage int
	doc := GetBody(url)
	pages, ok := doc.Find("div.pagination em").Attr("data-total-pages")
	if ok {
		totalPage, _ = strconv.Atoi(pages)
	}
	return totalPage
}
