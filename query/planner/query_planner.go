package planner

import (
	"github.com/pin-yu/miniature-db/query/algebra"
	"github.com/pin-yu/miniature-db/query/parser"
)

type QueryPlanner interface {
	CreatePlan(parser.QueryData) algebra.Plan
}
