package helper

type Where interface {
	Add(query interface{}, elem ...interface{}) Where
	GetQuery() interface{}
	GetArgs() []interface{}
	Next() bool
	Rewind()
}

type condition struct {
	query interface{}
	args  []interface{}
}

type where struct {
	index int
	con   []*condition
}

func (w *where) GetArgs() []interface{} {
	return w.con[w.index].args
}

func (w *where) GetQuery() interface{} {
	return w.con[w.index].query
}

func (w *where) Next() bool {
	w.index++
	return w.index < len(w.con)
}

func (w *where) Add(query interface{}, elem ...interface{}) Where {
	w.con = append(w.con, &condition{
		query: query,
		args:  elem,
	})
	return w
}

func (w *where) Rewind() {
	w.index = -1
}

func NewWhere() Where {
	return &where{
		index: -1,
		con:   []*condition{},
	}
}
