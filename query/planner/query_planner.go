package planner

import (
	"github.com/pin-yu/miniature-db/query/algebra"
	"github.com/pin-yu/miniature-db/query/parser"
)

type QueryPlanner interface {
	CreatePlan(parser.QueryData) algebra.Plan
}

type BasicQPlanner struct {
}

func (p *BasicQPlanner) CreatePlan(parser.QueryData) algebra.Plan {
	return &algebra.SelectPlan{}
}
