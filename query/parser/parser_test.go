package parser

import (
	"testing"

	"github.com/pin-yu/miniature-db/sqltype"
	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  string
		Expect QueryData
		Error  error
	}{
		{
			"1 projection 1 table",
			"SELECT c1 FROM t1;",
			QueryData{
				ProjFields: map[string]any{
					"c1": nil,
				},
				Tables: map[string]any{
					"t1": nil,
				},
				Predicate: sqltype.Predicate{},
			},
			nil,
		},
	}

	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.Name, func(t *testing.T) {
			p := NewParser(tc.Input)
			actual, err := p.ParseQuery()
			assert.Equal(t, tc.Expect, *actual)
			assert.NoError(t, err)
		})
	}
}
