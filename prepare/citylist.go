package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error : ", resp.StatusCode)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("%s\n", all)
	//PrintCityList(all)
	PrintCityListForJs(all)
}

func PrintCityListForJs(contents []byte)  {
	re := regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)
	all := re.FindSubmatch(contents)
	fmt.Printf("%s",all[1])

}
func PrintCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)
	for _, c := range all {
		fmt.Printf("city\t:%s, url\t:%s\n", c[2], c[1])
	}
}
