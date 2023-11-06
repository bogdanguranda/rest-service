package service

type Searcher interface {
	Search(value int) (int, int)
}
