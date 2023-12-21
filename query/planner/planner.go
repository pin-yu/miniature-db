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

func (p *Planner) CreateQueryPlan(qry string) (algebra.Plan, error) {
	parser := parser.NewParser(qry)
	queryData, err := parser.ParseQuery()
	if err != nil {
		return nil, err
	}

	// TODO: verify queryData before creating a plan

	return p.qPlanner.CreatePlan(queryData), nil
}
