package parser

import "fmt"

type UpdateType int

const (
	CREATE_TABLE UpdateType = iota
	INSERT
)

type Parser struct {
	l        *lexer
	curToken token
}

func NewParser(input string) *Parser {
	return &Parser{
		l: newLexer(input),
	}
}

func (p *Parser) ParseQuery() (*QueryData, error) {
	p.nextToken()
	if p.curToken.Lit != "select" {
		return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
	}

	projFields, err := p.parseProjFields()
	if err != nil {
		return nil, err
	}
	if len(projFields) == 0 {
		return nil, fmt.Errorf("project field is empty")
	}

	tables, err := p.parseTables()
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		return nil, fmt.Errorf("table field is empty")
	}

	return &QueryData{
		ProjFields: projFields,
		Tables:     tables,
	}, nil
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

func (p *Parser) nextToken() {
	p.curToken = p.l.NextToken()
}

func (p *Parser) parseProjFields() (map[string]struct{}, error) {
	projFields := map[string]struct{}{}
	lastComma := false
	for {
		p.nextToken()
		if p.curToken.Lit == "," {
			if len(projFields) != 0 && !lastComma {
				lastComma = true
				continue
			} else {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			}
		}
		lastComma = false

		if p.curToken.Lit == "from" {
			break
		}
		if p.curToken.TokenType == KW {
			return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
		}
		if p.curToken.Lit != "\"" && p.curToken.TokenType == DELIM {
			return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
		}
		if p.curToken.Lit == "\"" {
			p.nextToken()
			if p.curToken.TokenType != ID {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			} else {
				projFields[p.curToken.Lit] = struct{}{}
			}

			p.nextToken()
			if p.curToken.Lit != "\"" {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			}
		} else {
			projFields[p.curToken.Lit] = struct{}{}
		}
	}
	return projFields, nil
}

func (p *Parser) parseTables() (map[string]struct{}, error) {
	tables := map[string]struct{}{}
	lastComma := false
	for {
		p.nextToken()
		if p.curToken.Lit == "," {
			if len(tables) != 0 || !lastComma {
				lastComma = true
				continue
			} else {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			}
		}
		lastComma = false
		if p.curToken.Lit == "where" || p.curToken.Lit == ";" {
			break
		}
		if p.curToken.TokenType == KW {
			return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
		}
		if p.curToken.Lit != "\"" && p.curToken.TokenType == DELIM {
			return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
		}
		if p.curToken.Lit == "\"" {
			p.nextToken()
			if p.curToken.TokenType != ID {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			} else {
				tables[p.curToken.Lit] = struct{}{}
			}

			p.nextToken()
			if p.curToken.Lit != "\"" {
				return nil, fmt.Errorf("bad syntax, pos=%d", p.curToken.EndPos)
			}
		} else {
			tables[p.curToken.Lit] = struct{}{}
		}
	}
	return tables, nil
}
