package helper

type Select interface {
	GetQuery() interface{}
	GetArgs() []interface{}
}

type field struct {
	query interface{}
	args  []interface{}
}

func (f *field) GetQuery() interface{} {
	return f.query
}

func (f *field) GetArgs() []interface{} {
	return f.args
}

func NewSelect(query interface{}, args ...interface{}) *field {
	return &field{
		query: query,
		args:  args,
	}
}
