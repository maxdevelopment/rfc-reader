package extractor

import (
	"fmt"
	"regexp"
	"sync"
	s "strings"
	"sort"
	"strconv"
)

var PopWordsQty int
var WordLength int

var Rep = report{
	result: make(map[string]int),
}
var DataCh = make(chan string)
var wordsCh = make(chan extractor)

type result struct {
	key   string
	value int
}

func sortMap(m map[string]int) []result {
	var res []result
	for k, v := range m {
		res = append(res, result{k, v})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].value > res[j].value
	})

	return res
}

type extractor struct {
	words []string
}

func (ex *extractor) prepare() bool {

	for _, word := range ex.words {
		Rep.result[s.ToLower(word)] ++
	}

	sc := sortMap(Rep.result)

	Rep.clearMap()

	for _, v := range sc[:PopWordsQty] {
		Rep.result[v.key] = v.value
	}

	return true
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
	rep := sortMap(r.result)
	for k, v := range rep {
		fmt.Printf("%d %s [ %d ]\n", k, v.key, v.value)
	}
}

func Run(wg *sync.WaitGroup) {
	go extract(wg)
	go func() {

		for body := range DataCh {
			r := regexp.MustCompile(`\w{` + strconv.Itoa(WordLength) + `,}`)
			matches := r.FindAllString(body, -1)
			wordsCh <- extractor{
				words: matches,
			}
		}
	}()
}

func extract(wg *sync.WaitGroup) {
	for ex := range wordsCh {
		if ex.prepare() {
			wg.Done()
		}
	}
}
