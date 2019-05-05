package fetcher

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//抓取超时时间
const timeout = 30 * time.Second

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
	Timeout: timeout,
}

/*
fetcher：根据url获取对应的数据
*/
func Fetch(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp, err := httpClient.Do(request)
	if err != nil || resp == nil {
		log.Print(err)
		return nil, err
	}

	defer func() {
		switch p := recover(); p.(type) {
		case nil:
			resp.Body.Close()
		default:
			log.Println(p)
			//panic(p)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code:%d  ", resp.StatusCode)
	}
	//如果页面传来的不是utf8，我们需要转为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	uft8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(uft8Reader)
}

// 读取编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
