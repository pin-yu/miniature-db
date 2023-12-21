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
				{Lit: "", EndPos: 0, TokenType: END},
				{Lit: "", EndPos: 0, TokenType: END},
			},
		},
		{
			Name: "edge case 2",
			Text: " ",
			Expecteds: []token{
				{Lit: "", EndPos: 1, TokenType: END},
				{Lit: "", EndPos: 1, TokenType: END},
			},
		},
		{
			Name: "edge case 3",
			Text: "👍",
			Expecteds: []token{
				{Lit: "👍", EndPos: 1, TokenType: UNK},
			},
		},
		{
			Name: "delim 1",
			Text: "  ,  ;  ",
			Expecteds: []token{
				{Lit: ",", EndPos: 3, TokenType: DELIM},
				{Lit: ";", EndPos: 6, TokenType: DELIM},
				{Lit: "", EndPos: 8, TokenType: END},
				{Lit: "", EndPos: 8, TokenType: END},
			},
		},
		{
			Name: "DML insert into table",
			Text: "INSERT INTO tbl (c1, c2, c3) values ('str1', 2, 3);",
			Expecteds: []token{
				{Lit: "insert", EndPos: 6, TokenType: KW},
				{Lit: "into", EndPos: 11, TokenType: KW},
				{Lit: "tbl", EndPos: 15, TokenType: ID},
				{Lit: "(", EndPos: 17, TokenType: DELIM},
				{Lit: "c1", EndPos: 19, TokenType: ID},
				{Lit: ",", EndPos: 20, TokenType: DELIM},
				{Lit: "c2", EndPos: 23, TokenType: ID},
				{Lit: ",", EndPos: 24, TokenType: DELIM},
				{Lit: "c3", EndPos: 27, TokenType: ID},
				{Lit: ")", EndPos: 28, TokenType: DELIM},
				{Lit: "values", EndPos: 35, TokenType: KW},
				{Lit: "(", EndPos: 37, TokenType: DELIM},
				{Lit: "'", EndPos: 38, TokenType: DELIM},
				{Lit: "str1", EndPos: 42, TokenType: ID},
				{Lit: "'", EndPos: 43, TokenType: DELIM},
				{Lit: ",", EndPos: 44, TokenType: DELIM},
				{Lit: "2", EndPos: 46, TokenType: ID},
				{Lit: ",", EndPos: 47, TokenType: DELIM},
				{Lit: "3", EndPos: 49, TokenType: ID},
				{Lit: ")", EndPos: 50, TokenType: DELIM},
				{Lit: ";", EndPos: 51, TokenType: DELIM},
				{Lit: "", EndPos: 51, TokenType: END},
				{Lit: "", EndPos: 51, TokenType: END},
			},
		},
		{
			Name: "DQL select 1",
			Text: "SELECT a, b FROM table_c WHERE c1='str1';",
			Expecteds: []token{
				{Lit: "select", EndPos: 6, TokenType: KW},
				{Lit: "a", EndPos: 8, TokenType: ID},
				{Lit: ",", EndPos: 9, TokenType: DELIM},
				{Lit: "b", EndPos: 11, TokenType: ID},
				{Lit: "from", EndPos: 16, TokenType: KW},
				{Lit: "table_c", EndPos: 24, TokenType: ID},
				{Lit: "where", EndPos: 30, TokenType: KW},
				{Lit: "c1", EndPos: 33, TokenType: ID},
				{Lit: "=", EndPos: 34, TokenType: DELIM},
				{Lit: "'", EndPos: 35, TokenType: DELIM},
				{Lit: "str1", EndPos: 39, TokenType: ID},
				{Lit: "'", EndPos: 40, TokenType: DELIM},
				{Lit: ";", EndPos: 41, TokenType: DELIM},
				{Lit: "", EndPos: 41, TokenType: END},
				{Lit: "", EndPos: 41, TokenType: END},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := newLexer(tc.Text)
			for _, tok := range tc.Expecteds {
				assert.Equal(t, tok, l.NextToken())
			}
		})
	}
}
