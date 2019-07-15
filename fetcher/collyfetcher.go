package fetcher

import (
	"fmt"
	"github.com/gocolly/colly"
)

// colly的抓取

//根据url获取对应的数据
func FetchColly() {
	// 声明初始化NewCollector对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等
	c := colly.NewCollector(
		//colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36"))

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", "news.baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Referer", "http://www.baidu.com")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		fmt.Println("Visiting", r.URL)
	})

	// 对响应的HTML元素处理
	c.OnHTML("title", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println("title:", e.Text)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		e.ForEach(".hotnews li strong a", func(i int, el *colly.HTMLElement) {
			band := el.Attr("href")
			title := el.Text
			fmt.Printf("新闻 %d : %s - %s\n", i, title, band)
			// e.Request.Visit(band)
		})
	})

	// 发现并访问下一个连接
	//c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
	//    e.Request.Visit(e.Attr("href"))
	//})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
		// 设置context
		// fmt.Println(r.Ctx.Get("url"))
		//fmt.Println(string(r.Body))
		// goquery直接读取resp.Body的内容
		//htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		//if err != nil {
		//	log.Fatal(err)
		//}
		//find := htmlDoc.Find(".hotnews li strong a")
		//
		//find.Each(func(i int, s *goquery.Selection) {
		//	band, _ := s.Attr("href")
		//	title := s.Text()
		//	fmt.Printf("热点新闻 %d: %s - %s\n", i, title, band)
		//	c.Visit(band)
		//})
	})

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	c.Visit("http://news.baidu.com")
}
