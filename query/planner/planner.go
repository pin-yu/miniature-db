package planner

import (
	"github.com/pin-yu/miniature-db/query/algebra"
	"github.com/pin-yu/miniature-db/query/parser"
)

var _ QueryPlanner = (*BasicQPlanner)(nil)

type Planner struct {
	qPlanner      QueryPlanner
	updatePlanner UpdatePlanner
}

func NewPlanner(qPlanner QueryPlanner, uPlanner UpdatePlanner) *Planner {
	return &Planner{
		qPlanner:      qPlanner,
		updatePlanner: uPlanner,
	}
}

func (p *Planner) CreateQueryPlan(qry string) algebra.Plan {
	parser := parser.NewParser(qry)
	queryData := parser.ParseQuery()

	// TODO: verify queryData before creating a plan

	return p.qPlanner.CreatePlan(queryData)
}
