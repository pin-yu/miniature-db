package parser

type UpdateType int

const (
	CREATE_TABLE UpdateType = iota
	INSERT
)

type Parser struct {
	l *lexer
}

func NewParser(input string) *Parser {
	return &Parser{
		l: newLexer(input),
	}
}

func (p *Parser) ParseQuery() QueryData {
	return QueryData{}
}

func (p *Parser) ParseUpdate() UpdateType {
	return CREATE_TABLE
}

func (p *Parser) ParseCreateTbl() CreateTblData {
	return CreateTblData{}
}

func (p *Parser) ParseInsert() InsertData {
	return InsertData{}
}
