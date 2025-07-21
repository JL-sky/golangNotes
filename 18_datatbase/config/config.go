package config

const (
	TOutput     = "t_output"
	TChangeLogs = "t_changelog"
)

var FilterFields = map[string]struct{}{
	"RecordStatus": {},
	"CMtime":       {},
}
