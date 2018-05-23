package parser

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"rfc-reader/extractor"
	"sync"
)

type webCrawler struct {
	host    string
	rpcQty  int
	connQty int
	links   chan string
}

func WebParams(rpcQty int, connQty int) *webCrawler {
	return &webCrawler{
		"https://tools.ietf.org/rfc/rfc",
		rpcQty,
		connQty,
		make(chan string),
	}
}

func fetch(url string) (content string, err error) {
	fmt.Println(url)
	resp, err := http.Get(url)

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

func newConn(crawler *webCrawler, wg *sync.WaitGroup) {
	fmt.Println("CONN: ")
	wg.Add(1)
	defer wg.Done()

	for link := range crawler.links {

		body, err := fetch(link)
		if err != nil {
			fmt.Errorf("parser error link: %s | error: %v", link, err)
			return
		}

		extractor.DataCh <- body
	}
}

func (crawler *webCrawler) parse() {
	wg := new(sync.WaitGroup)
	for i := 0; i <= crawler.connQty; i++ {
		fmt.Println("CONN: ", i)
		go newConn(crawler, wg)
		//go func() {
		//	wg.Add(1)
		//	defer wg.Done()
		//	for link := range crawler.links {
		//
		//		body, err := fetch(link)
		//		if err != nil {
		//			fmt.Errorf("parser error link: %s | error: %v", link, err)
		//			return
		//		}
		//
		//		extractor.DataCh <- body
		//	}
		//}()
	}
	wg.Wait()

	for i := 1; i <= crawler.rpcQty; i++ {
		fmt.Println("num: ", i)
		link := crawler.host + strconv.Itoa(i) + ".txt"
		crawler.links <- link
	}
}
