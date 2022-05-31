package helper

type Group interface {
	GetGroup() string
}

type group struct {
	statment string
}

func (g *group) GetGroup() string {
	return g.statment
}

func NewGroup(statment string) *group {
	return &group{
		statment: statment,
	}
}

type Having interface {
	GetQuery() interface{}
	GetArgs() []interface{}
}

type having struct {
	query interface{}
	args  []interface{}
}

func (h *having) GetQuery() interface{} {
	return h.query
}

func (h *having) GetArgs() []interface{} {
	return h.args
}

func NewHaving(query interface{}, args ...interface{}) *having {
	return &having{
		query: query,
		args:  args,
	}
}
