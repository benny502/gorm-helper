package builder

import (
	"gorm.io/gorm"
)

type options struct {
	where     Where
	associate Associate
}

type Builder interface {
	WithWhere(where Where) Builder
	WithAssociate(associate Associate) Builder
	Build(db *gorm.DB) *gorm.DB
}

type builder struct {
	opts options
}

func (b builder) WithWhere(where Where) Builder {
	b.opts.where = where
	return &b
}

func (b builder) WithAssociate(associate Associate) Builder {
	b.opts.associate = associate
	return &b
}

func (b *builder) Build(db *gorm.DB) *gorm.DB {
	tx := db.Assign()
	if b.opts.where != nil {
		for b.opts.where.Next() {
			args := b.opts.where.GetArgs()
			if len(args) > 0 {
				tx.Where(b.opts.where.GetQuery(), args...)
			} else {
				tx.Where(b.opts.where.GetQuery())
			}
		}
	}

	if b.opts.associate != nil {
		tx.Preload(b.opts.associate.GetPreload())
		joins := b.opts.associate.GetJoinsString()
		for _, join := range joins {
			tx.Joins(join)
		}
	}
	return tx
}

func NewBuilder() Builder {
	return &builder{}
}

type Where interface {
	Add(query interface{}, elem ...interface{}) Where
	GetQuery() interface{}
	GetArgs() []interface{}
	Next() bool
}

type Associate interface {
	GetPreload() string
	GetJoinsString() []string
}
