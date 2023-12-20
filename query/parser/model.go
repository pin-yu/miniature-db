package parser

import "github.com/pin-yu/miniature-db/sqltype"

type InsertData struct {
}

type CreateTblData struct {
	TblName string
	Schema  sqltype.Schema
}

type QueryData struct {
}
