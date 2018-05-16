package main

import (
	"flag"
	"fmt"
)

const (
	defaultRrcQuantity         = 1000
	defaultRfcQuantityHelper   = "quantity RFC for wrapping"
	popularWordsQuantity       = 20
	popularWordsQuantityHelper = "popular words quantity"
	popularWordLength          = 4
	popularWordLengthHelper    = "popular words minimal length"
	connectionsQuantity        = 10
	connectionsQuantityHelper  = "connections quantity"
)

func main() {
	rpcQty := flag.Int("rpcQty", defaultRrcQuantity, defaultRfcQuantityHelper)
	popWordsQty := flag.Int("wordsQty", popularWordsQuantity, popularWordsQuantityHelper)
	wordLengh := flag.Int("wordLen", popularWordLength, popularWordLengthHelper)
	connQty := flag.Int("connQty", connectionsQuantity, connectionsQuantityHelper)

	flag.Parse()
	fmt.Println(*rpcQty)
	fmt.Println(*popWordsQty)
	fmt.Println(*wordLengh)
	fmt.Println(*connQty)
}
