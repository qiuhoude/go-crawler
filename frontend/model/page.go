package model

type SearchResult struct {
	Hits     int64 //所有的数据量
	Start    int   //从多少开始
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{} //展示的每条用户信息
}
