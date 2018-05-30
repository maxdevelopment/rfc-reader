package parser

import "sync"

type Parser interface {
	parse(wg *sync.WaitGroup)
}

func Start(parser Parser, wg *sync.WaitGroup) {
	parser.parse(wg)
}
