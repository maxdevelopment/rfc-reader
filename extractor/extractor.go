package extractor

import (
	"fmt"
	"regexp"
	"sync"
)

var DataCh = make(chan string)

func Run(wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {

		for body := range DataCh {
			fmt.Println("RECEIVED")

			r := regexp.MustCompile(`\w{4,}`)
			matches := r.FindAllString(body, -1)
			fmt.Println(len(matches))
			fmt.Println(matches)

			wg.Done()

		}
	}(wg)

}
