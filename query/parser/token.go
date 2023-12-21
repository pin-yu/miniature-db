package parser

type TokenType int

const (
	KW TokenType = iota
	ID
	DELIM

	END
	UNK
)

type token struct {
	Lit       string
	EndPos    int
	TokenType TokenType
}
