package planner

import "github.com/pin-yu/miniature-db/query/parser"

type UpdatePlanner interface {
	ExecuteInsert(parser.InsertData) (int, error)
	ExecuteCreateTbl(parser.CreateTblData) (int, error)
}
