package controller

import (
	"context"
	"crawler/engine"
	"crawler/frontend/model"
	"crawler/frontend/view"
	"crawler/parser"
	"crawler/persist"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(persist.EsUrl))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	log.Printf("q:%s, from:%d\n", q, from)

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.view.Render(w, page)
	log.Println("...", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	result.Query = q

	resp, err := h.client.Search(parser.ES_INDEX).
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	log.Println("resp-->", resp)
	result.Hits = resp.TotalHits()
	result.Start = from

	itemRaw := resp.Each(reflect.TypeOf(engine.Item{}))
	for _, v := range itemRaw {
		log.Printf("%+v\n", v)
		item, ok := v.(engine.Item)
		if ok && item.PayLoad != nil {
			result.Items = append(result.Items, item)
		}
	}
	return result, nil
}
