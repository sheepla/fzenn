package main

import (
	"fmt"
	"log"

	"github.com/sheepla/fzenn/client"
)

func main() {
	a, err := client.FetchArticleList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Articles[0].Title)

	u, err := client.FetchUser("sheepla")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", u.User.Username, u.User.Bio)

	tl, err := client.FetchTopicsList()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range tl.Topics {
		fmt.Printf("%s #%s count: %d\n", v.DisplayName, v.Name, v.TaggingsCount)
	}

	t, err := client.FetchTopics("機械学習")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range t.Articles {
		fmt.Printf("%s\n", v.Title)
	}
}
