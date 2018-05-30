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
	key string
	value int
}

var res []result
var m = make(map[string]int)

func (ex *extractor) sort() bool {

	//m := make(map[string]int)

	for _, word := range ex.words {
		m[s.ToLower(word)] ++
	}

	//type kv struct {
	//	Key   string
	//	Value int
	//}
	//
	//var ss []kv
	for k, v := range m {
		res = append(res, result{k, v})
	}
	fmt.Println(res)

	sort.Slice(res, func(i, j int) bool {
		return res[i].value > res[j].value
	})

	//fmt.Println(res)
	res := res[:10]
	fmt.Println(res)


	//fmt.Println("PRINT")
	//for _, kv := range res {
	//	fmt.Printf("%s, %d\n", kv.key, kv.value)
	//	Result[kv.key] = kv.value
	//}

	return true
}

var Result = make(map[string]int)
var DataCh = make(chan string)
var wordsCh = make(chan extractor)

func Run(wg *sync.WaitGroup) {
	go extract(wg)
	go func() {

		for body := range DataCh {
			fmt.Println("RECEIVED")

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
