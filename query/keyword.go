package query

func init() {
	for _, s := range kws {
		kwSet[s] = true
	}

	for _, d := range delims {
		delimSet[d] = true
	}
}

var kws = []string{
	"database",
	"table",
	"column",

	// Data Definition Language
	"create",
	"alter",
	"add",
	"drop",
	"truncate",

	// Data Manipulation Language
	"insert",
	"into",
	"values",
	"update",
	"set",
	"delete",

	// Data Query Language
	"select",
	"from",

	// Data Control Language
	"commit",
	"rollback",

	// condition
	"where",
}

var kwSet = map[string]bool{}

func isKw(s string) bool {
	_, exist := kwSet[s]
	return exist
}

var delims = []rune{
	',',
	';',
	'(',
	')',
	'\'',
	'=',
	' ',
}

var delimSet = map[rune]bool{}

func isDelim(r rune) bool {
	_, exist := delimSet[r]
	return exist
}
