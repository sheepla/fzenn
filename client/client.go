package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func newURL(urlPath string) string {
	u := &url.URL{}
	u.Scheme = "https"
	u.Host = "zenn-api.deno.dev"
	if urlPath != "" {
		u.Path = urlPath
	}
	return u.String()
}

func fetch(url string) (*[]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func FetchTrending() (*Trending, error) {
	b, err := fetch(newURL(""))
	if err != nil {
		return nil, err
	}

	var t Trending
	if err := json.Unmarshal(*b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func FetchArticleList() (*ArticleList, error) {
	b, err := fetch(newURL("articles"))
	if err != nil {
		return nil, err
	}

	var a ArticleList
	if err := json.Unmarshal(*b, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func FetchUser(name string) (*User, error) {
	b, err := fetch(newURL(name))
	if err != nil {
		return nil, err
	}

	var u User
	if err := json.Unmarshal(*b, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func FetchTopicsList() (*TopicsList, error) {
	b, err := fetch(newURL("topics"))
	if err != nil {
		return nil, err
	}

	var t TopicsList
	if err := json.Unmarshal(*b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func FetchTopics(name string) (*Topics, error) {
	b, err := fetch(newURL(path.Join("topics", name)))
	if err != nil {
		return nil, err
	}

	var t Topics
	if err := json.Unmarshal(*b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
