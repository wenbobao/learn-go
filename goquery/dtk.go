package main

import (
	"fmt"
	"log"
	"net/http"
  "github.com/PuerkitoBio/goquery"
  "strings"
)

func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("http://www.dataoke.com/top_sell?cid=4&type=1")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find(".goods-list ").Children().EachWithBreak(func(i int, s *goquery.Selection) bool {
    // 打印
    fmt.Println(s.Html())
    // 商品标题
    title := s.Find(".goods-tit-txt a").Text()
    title = strings.TrimSpace(title)
    title = strings.Replace(title, " ", "", -1)
    title = strings.Replace(title, "\n", "", -1)
    // 图片URL
    items := s.Find("img")
    fmt.Println(items.Html())
    // goods_image := s.Find(".goods-img").Find("a").Find("img.img").Eq(1).Attr("data-original")
    // 商品2小时销量
    twohoursale := s.Find(".tg-show b").Text()
    // 商品券后价
    // goods_price := s.Find(".goods-img a").Attr("data-config")
    // 商品🈷月销量
    goods_month_sale := s.Find(".goods-sale b").Text()

    fmt.Printf("Review %d: %s - %s - %s - %s\n", i, title, twohoursale, goods_month_sale, goods_month_sale)

    if i == 0 {
      return false
    }
    return true
  })
}

func main() {
  ExampleScrape()
}

