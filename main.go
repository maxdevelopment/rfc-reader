package main

import (
	"github.com/maxdevelopment/rfc-reader/parser"
	"fmt"
	"errors"
	"flag"
	"os"
	"github.com/maxdevelopment/rfc-reader/extractor"
	"sync"
)

const (
	defaultRfcQuantity       = 3
	defaultRfcQuantityHelper = "quantity RFC for wrapping"
	//popularWordsQuantity       = 20
	//popularWordsQuantityHelper = "popular words quantity"
	//popularWordLength          = 4
	//popularWordLengthHelper    = "popular words minimal length"
	connectionsQuantity       = 4
	connectionsQuantityHelper = "connections quantity"
	parserSource              = "web"
	parserSourceHelper        = "parser source(web|files|db)"
)

func getType() (parser.Parser, error) {
	rpcQty := flag.Int("rpcQty", defaultRfcQuantity, defaultRfcQuantityHelper)
	////popWordsQty := flag.Int("wordsQty", popularWordsQuantity, popularWordsQuantityHelper)
	////wordLength := flag.Int("wordLen", popularWordLength, popularWordLengthHelper)
	connQty := flag.Int("connQty", connectionsQuantity, connectionsQuantityHelper)
	parserSrc := flag.String("src", parserSource, parserSourceHelper)
	flag.Parse()

	switch *parserSrc {
	case "web":
		webCrawler := parser.WebParams(*rpcQty, *connQty)
		return webCrawler, nil
	default:
		return nil, errors.New("source not implemented")
	}
}

func main() {

	crawlerType, err := getType()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	wg := new(sync.WaitGroup)
	wg.Add(defaultRfcQuantity)

	extractor.Run(wg)

	parser.Start(crawlerType, wg)

	wg.Wait()
	defer func() {
		fmt.Println("ENDED")
		fmt.Println(extractor.Result)
		//61,27,27,26,24
	}()
}
