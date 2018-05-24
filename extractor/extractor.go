package extractor

import (
	"fmt"
	"regexp"
)

var DataCh = make(chan string)

func Run() {
	fmt.Println("Extractor RUN")
	for {
		select {
		case dataCh := <-DataCh:

			r := regexp.MustCompile(`\w{4,}`)
			matches := r.FindAllString(dataCh, -1)
			fmt.Println(len(matches))
			fmt.Println(matches)

		//case doneCh := <-DoneCh:
		//	if doneCh {
		//		close(DataCh)
		//		close(DoneCh)
		//	}
		//default:
		//	fmt.Println("nothing ready")
		}
	}
}
