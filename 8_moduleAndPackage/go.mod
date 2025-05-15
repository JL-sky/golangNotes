module github.com/moduleExp

replace (
	github.com/string_utils => ./utils
	github.com/test => ./test
)

go 1.24.2

require (
	github.com/string_utils v0.0.0-00010101000000-000000000000
	github.com/test v0.0.0-00010101000000-000000000000
)
