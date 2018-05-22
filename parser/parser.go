package parser

type Parser interface {
	parse()
}

func Start(parser Parser) {
	parser.parse()
}
