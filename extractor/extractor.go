package extractor

import (
	"fmt"
	"regexp"
)

var DataCh = make(chan string)
var DoneCh = make(chan bool)

func Run() {
	for {
		select {
		case dataCh := <-DataCh:

			r := regexp.MustCompile(`\w{4,}`)
			matches := r.FindAllString(dataCh, -1)
			fmt.Println(matches)

		case doneCh := <-DoneCh:
			fmt.Println(doneCh)
		}
	}
}
