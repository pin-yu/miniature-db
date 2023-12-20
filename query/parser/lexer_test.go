package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		Name      string
		Text      string
		Expecteds []token
	}{
		{
			Name: "edge case 1",
			Text: "",
			Expecteds: []token{
				{Lit: "", TokenType: END},
				{Lit: "", TokenType: END},
			},
		},
		{
			Name: "edge case 2",
			Text: " ",
			Expecteds: []token{
				{Lit: "", TokenType: END},
				{Lit: "", TokenType: END},
			},
		},
		{
			Name: "edge case 3",
			Text: "👍",
			Expecteds: []token{
				{Lit: "👍", TokenType: UNK},
			},
		},
		{
			Name: "delim 1",
			Text: "  ,  ;  ",
			Expecteds: []token{
				{Lit: ",", TokenType: DELIM},
				{Lit: ";", TokenType: DELIM},
				{Lit: "", TokenType: END},
				{Lit: "", TokenType: END},
			},
		},
		{
			Name: "DML insert into table",
			Text: "INSERT INTO tbl (c1, c2, c3) values ('str1', 2, 3);",
			Expecteds: []token{
				{Lit: "insert", TokenType: KW},
				{Lit: "into", TokenType: KW},
				{Lit: "tbl", TokenType: ID},
				{Lit: "(", TokenType: DELIM},
				{Lit: "c1", TokenType: ID},
				{Lit: ",", TokenType: DELIM},
				{Lit: "c2", TokenType: ID},
				{Lit: ",", TokenType: DELIM},
				{Lit: "c3", TokenType: ID},
				{Lit: ")", TokenType: DELIM},
				{Lit: "values", TokenType: KW},
				{Lit: "(", TokenType: DELIM},
				{Lit: "'", TokenType: DELIM},
				{Lit: "str1", TokenType: ID},
				{Lit: "'", TokenType: DELIM},
				{Lit: ",", TokenType: DELIM},
				{Lit: "2", TokenType: ID},
				{Lit: ",", TokenType: DELIM},
				{Lit: "3", TokenType: ID},
				{Lit: ")", TokenType: DELIM},
				{Lit: ";", TokenType: DELIM},
				{Lit: "", TokenType: END},
				{Lit: "", TokenType: END},
			},
		},
		{
			Name: "DQL select 1",
			Text: "SELECT a, b FROM table_c WHERE c1='str1';",
			Expecteds: []token{
				{Lit: "select", TokenType: KW},
				{Lit: "a", TokenType: ID},
				{Lit: ",", TokenType: DELIM},
				{Lit: "b", TokenType: ID},
				{Lit: "from", TokenType: KW},
				{Lit: "table_c", TokenType: ID},
				{Lit: "where", TokenType: KW},
				{Lit: "c1", TokenType: ID},
				{Lit: "=", TokenType: DELIM},
				{Lit: "'", TokenType: DELIM},
				{Lit: "str1", TokenType: ID},
				{Lit: "'", TokenType: DELIM},
				{Lit: ";", TokenType: DELIM},
				{Lit: "", TokenType: END},
				{Lit: "", TokenType: END},
			},
		},
	}

	for _, tc := range testCases {
		l := newLexer(tc.Text)
		for _, tok := range tc.Expecteds {
			assert.Equal(t, tok, l.NextToken())
		}
	}
}
