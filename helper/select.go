package helper

type Select interface {
	GetQuery() interface{}
	GetArgs() []interface{}
}

type field struct {
	query interface{}
	args  []interface{}
}

func NewSelect(query interface{}, args ...interface{}) *field {
	return &field{
		query: query,
		args:  args,
	}
}
