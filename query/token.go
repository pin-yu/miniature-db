package query

type TokenType int

const (
	KW TokenType = iota
	ID
	DELIM

	END
	UNK
)

type Token struct {
	Lit       string
	TokenType TokenType
}
