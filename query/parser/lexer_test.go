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
		// {
		// 	Name: "DDL create table",
		// 	Text: "CREATE TABLE tbl;",
		// 	Expecteds: []token{
		// 		{Lit: "create", TokenType: KW},
		// 		{Lit: "table", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
		// {
		// 	Name: "DDL alter table",
		// 	Text: "ALTER TABLE tbl ADD c1 text;",
		// 	Expecteds: []token{
		// 		{Lit: "alter", TokenType: KW},
		// 		{Lit: "table", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: "add", TokenType: KW},
		// 		{Lit: "c1", TokenType: ID},
		// 		{Lit: "text", TokenType: ID},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
		// {
		// 	Name: "DDL drop table",
		// 	Text: "DROP TABLE tbl;",
		// 	Expecteds: []token{
		// 		{Lit: "drop", TokenType: KW},
		// 		{Lit: "table", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
		// {
		// 	Name: "DDL truncate table",
		// 	Text: "TRUNCATE TABLE tbl;",
		// 	Expecteds: []token{
		// 		{Lit: "truncate", TokenType: KW},
		// 		{Lit: "table", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
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
		// {
		// 	Name: "DML update rows",
		// 	Text: "UPDATE tbl SET c1 = 'str2', c2 = 4, c3 = 5 WHERE c1 = 'str1';",
		// 	Expecteds: []token{
		// 		{Lit: "update", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: "set", TokenType: KW},
		// 		{Lit: "c1", TokenType: ID},
		// 		{Lit: "=", TokenType: DELIM},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: "str2", TokenType: ID},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: ",", TokenType: DELIM},
		// 		{Lit: "c2", TokenType: ID},
		// 		{Lit: "=", TokenType: DELIM},
		// 		{Lit: "4", TokenType: ID},
		// 		{Lit: ",", TokenType: DELIM},
		// 		{Lit: "c3", TokenType: ID},
		// 		{Lit: "=", TokenType: DELIM},
		// 		{Lit: "5", TokenType: ID},
		// 		{Lit: "where", TokenType: KW},
		// 		{Lit: "c1", TokenType: ID},
		// 		{Lit: "=", TokenType: DELIM},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: "str1", TokenType: ID},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
		// {
		// 	Name: "DML delete rows",
		// 	Text: "DELETE FROM tbl WHERE c1 = 'str1';",
		// 	Expecteds: []token{
		// 		{Lit: "delete", TokenType: KW},
		// 		{Lit: "from", TokenType: KW},
		// 		{Lit: "tbl", TokenType: ID},
		// 		{Lit: "where", TokenType: KW},
		// 		{Lit: "c1", TokenType: ID},
		// 		{Lit: "=", TokenType: DELIM},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: "str1", TokenType: ID},
		// 		{Lit: "'", TokenType: DELIM},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
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
		// {
		// 	Name: "DCL commit",
		// 	Text: "COMMIT;",
		// 	Expecteds: []token{
		// 		{Lit: "commit", TokenType: KW},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
		// {
		// 	Name: "DCL rollback",
		// 	Text: "ROLLBACK;",
		// 	Expecteds: []token{
		// 		{Lit: "rollback", TokenType: KW},
		// 		{Lit: ";", TokenType: DELIM},
		// 		{Lit: "", TokenType: END},
		// 		{Lit: "", TokenType: END},
		// 	},
		// },
	}

	for _, tc := range testCases {
		l := NewLexer(tc.Text)
		for _, tok := range tc.Expecteds {
			assert.Equal(t, tok, l.NextToken())
		}
	}
}