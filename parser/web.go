package parser

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"rfc-reader/extractor"
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

func (crawler *webCrawler) parse() {

	for i := 0; i <= crawler.connQty; i++ {
		go func() {
			for link := range crawler.links {
				fmt.Println(link)

				resp, err := http.Get(link)
				if resp != nil {
					defer resp.Body.Close()
				}

				if err != nil {
					fmt.Println(err)
					return
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				extractor.DataCh <- string(body)
			}
		}()
	}

	for i := 1; i <= crawler.rpcQty; i++ {
		link := crawler.host + strconv.Itoa(i) + ".txt"
		crawler.links <- link
	}
}
