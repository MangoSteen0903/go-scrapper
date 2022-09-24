package utils

import "fmt"

var repositoryList []RepositoryRequest
var urls []string
var result = make(map[string]RepositoryRequest)

func ScrapURL(baseURL string) []RepositoryRequest {
	repositoryCh := make(chan []RepositoryRequest)
	totalPage := GetPage(baseURL)

	for i := 1; i <= totalPage; i++ {
		currentPage := fmt.Sprintf("?page=%d", i)
		currentURL := baseURL + currentPage

		go GetRepositories(currentURL, repositoryCh)
	}

	for i := 0; i < totalPage; i++ {
		repository := <-repositoryCh
		repositoryList = append(repositoryList, repository...)
	}

	return repositoryList
}
