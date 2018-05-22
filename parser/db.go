package parser

import "fmt"

type dbCrawler struct {}

func (crawler *dbCrawler) parse() {
	fmt.Println("DB Parce() called")
}
