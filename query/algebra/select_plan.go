package algebra

import "github.com/pin-yu/miniature-db/sqltype"

type SelectPlan struct {
}

func (p *SelectPlan) Open() Scan {
	return &SelectScan{}
}

func (p *SelectPlan) Schema() *sqltype.Schema {
	return sqltype.NewSchema()
}
func (p *SelectPlan) RecordsOutput() int64 {
	return 0
}
