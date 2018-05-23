package main

import (
	"rfc-reader/parser"
	"fmt"
	"errors"
	"flag"
	"os"
	"rfc-reader/extractor"
)

const (
	defaultRfcQuantity       = 100
	defaultRfcQuantityHelper = "quantity RFC for wrapping"
	//popularWordsQuantity       = 20
	//popularWordsQuantityHelper = "popular words quantity"
	//popularWordLength          = 4
	//popularWordLengthHelper    = "popular words minimal length"
	connectionsQuantity       = 10
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

	go extractor.Run()
	parser.Start(crawlerType)

	var input string
	fmt.Scanln(&input)
}
