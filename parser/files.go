package parser

import (
	"fmt"
	"sync"
)

type filesCrawler struct {}

func (crawler *filesCrawler) parse(wg *sync.WaitGroup) {
	fmt.Println("FILE Parce() called")
}
