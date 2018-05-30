package extractor

import (
	"fmt"
	"regexp"
	"sync"
	s "strings"
)

type extractor struct {
	words []string
}

func (ex *extractor) sort() bool {
	for _, word := range ex.words {
		result[s.ToLower(word)] ++
	}

	return true
}

var result = make(map[string]int)
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
