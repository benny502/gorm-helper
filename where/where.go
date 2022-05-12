package where

import "github.com/benny502/gorm-helper/builder"

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

func (w *where) Add(query interface{}, elem ...interface{}) builder.Where {
	w.con = append(w.con, &condition{
		query: query,
		args:  elem,
	})
	return w
}

func NewWhere() builder.Where {
	return &where{
		index: -1,
		con:   []*condition{},
	}
}
