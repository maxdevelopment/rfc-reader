package extractor

import (
	"fmt"
	"regexp"
	"sync"
	s "strings"
	"sort"
)

type extractor struct {
	words []string
}

type result struct {
	key   string
	value int
}

type report struct {
	result map[string]int
}

func (r *report) clearMap() {
	for key := range r.result {
		delete(r.result, key)
	}
}

func (r *report) GetReport() {
	rep := sortAndCut(r.result)
	for k, v := range rep {
		fmt.Printf("%d %s [ %d ]\n", k, v.key, v.value)
	}
}

var Rep = report{
	result: make(map[string]int),
}

func sortAndCut(m map[string]int) []result {
	var res []result
	for k, v := range m {
		res = append(res, result{k, v})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].value > res[j].value
	})

	return res
}

func (ex *extractor) sort() bool {

	for _, word := range ex.words {
		Rep.result[s.ToLower(word)] ++
	}

	sc := sortAndCut(Rep.result)

	Rep.clearMap()

	for _, v := range sc[:20] {
		Rep.result[v.key] = v.value
	}

	return true
}

var DataCh = make(chan string)
var wordsCh = make(chan extractor)

func Run(wg *sync.WaitGroup) {
	go extract(wg)
	go func() {

		for body := range DataCh {
			r := regexp.MustCompile(`\w{4,}`)
			matches := r.FindAllString(body, -1)
			wordsCh <- extractor{
				words: matches,
			}
		}
	}()
}

func extract(wg *sync.WaitGroup) {
	for ex := range wordsCh {
		if ex.sort() {
			wg.Done()
		}
	}
}
