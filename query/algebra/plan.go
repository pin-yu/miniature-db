package algebra

import "github.com/pin-yu/miniature-db/sqltype"

type Plan interface {
	Open() Scan
	Schema() *sqltype.Schema
	RecordsOutput() int64
}
