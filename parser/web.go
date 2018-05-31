package parser

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"sync"
	"github.com/maxdevelopment/rfc-reader/extractor"
)

type webCrawler struct {
	host    string
	rpcQty  int
	connQty int
	links   chan string
}

func WebParams(rpcQty, connQty int) *webCrawler {
	return &webCrawler{
		"https://tools.ietf.org/rfc/rfc",
		rpcQty,
		connQty,
		make(chan string),
	}
}

func fetch(link string) (content string, err error) {
	resp, err := http.Get(link)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (crawler *webCrawler) parse(wg *sync.WaitGroup) {

	wg.Add(crawler.connQty)

	go func() {
		for i := 1; i <= crawler.rpcQty; i++ {
			link := crawler.host + strconv.Itoa(i) + ".txt"
			crawler.links <- link
		}
		close(crawler.links)
	}()

	for i := 0; i < crawler.connQty; i++ {
		go worker(i, crawler, wg)
	}
}

func worker(i int, crawler *webCrawler, wg *sync.WaitGroup) {
	for link := range crawler.links {

		fmt.Println("worker", i, "recieved", link)

		body, err := fetch(link)
		if err != nil {
			fmt.Errorf("parser error link: %s | error: %v", link, err)
			return
		}

		extractor.DataCh <- body

	}
	fmt.Println("worker", i, "done")
	wg.Done()
}
