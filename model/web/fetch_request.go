package web

type FetchRequest struct {
	Keyword string `schema:"keyword"`
	Sort 	string `schema:"sort"`
}
