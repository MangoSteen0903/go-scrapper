package utils

import (
	"log"
	"strings"
)

func CleanBadgeString(str string) []string {
	s := replacer.Replace(strings.TrimSpace(str))
	result := strings.Split(s, " ")
	return result
}

func CleanString(str string) string {
	s := strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
	return s
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
