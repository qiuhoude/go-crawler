package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//http.Handle("/", http.FileServer("frontend/view"))

	//http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err)
	//}

	content, err := ioutil.ReadFile("view/template_test.html")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", content)
	}
}
