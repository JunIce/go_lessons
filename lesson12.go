package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Spider struct {
	url    string
	header map[string]string
}

func (spider Spider) fetch() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", spider.url, nil)
	if err != nil {
	}

	for key, value := range spider.header {
		req.Header.Add(key, value)
	}
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func main() {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	spider := &Spider{"https://movie.douban.com/top250", header}
	body := spider.fetch()
	fmt.Println(body)
}
