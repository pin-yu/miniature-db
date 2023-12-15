package parser

type Parser struct {
	l *lexer
}

func NewParser(input string) *Parser {
	return &Parser{
		l: NewLexer(input),
	}
}

func ParseQuery() {

}
