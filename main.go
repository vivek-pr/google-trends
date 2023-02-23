package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel  `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Item  []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Traffic     string `xml:"ht:approx_traffic"`
	News		[]News `xml:"ht:news_item"`
}

type News struct {
	Headline string `xml:"ht:news_item_title"`
	HeadlineLink string `xml:"ht:news_item_url"`

}

func main() {
	var r RSS

	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for i, item := range r.Channel.Item {
		rank := i + 1
		fmt.Println("Rank", rank)
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Traffic)
		for _, news := range item.News {
			
			fmt.Println(news.Headline)
			fmt.Println(news.HeadlineLink)
		}
	}
}

func readGoogleTrends() [] byte{
	resp := getGoogleTrends()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return body
}

func getGoogleTrends() *http.Response{
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=IN")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}