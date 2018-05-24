package parser

import (
	"fmt"
	"sync"
)

type dbCrawler struct {}

func (crawler *dbCrawler) parse(wg *sync.WaitGroup) {
	fmt.Println("DB Parce() called")
}
