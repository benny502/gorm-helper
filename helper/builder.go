package helper

import (
	"github.com/benny502/gorm-helper/associate"
	"gorm.io/gorm"
)

type Builder interface {
	WithWhere(where Where) Builder
	WithAssociate(associate associate.Associate) Builder
	Build(db *gorm.DB) *gorm.DB
}

type options struct {
	where     Where
	associate associate.Associate
	preload   associate.Preload
}

type builder struct {
	opts options
}

func (b builder) WithWhere(where Where) Builder {
	b.opts.where = where
	return &b
}

func (b builder) WithAssociate(associate associate.Associate) Builder {
	b.opts.associate = associate
	return &b
}

func (b builder) WithPreload(preload associate.Preload) Builder {
	b.opts.preload = preload
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

	if b.opts.preload != nil {
		tx.Preload(b.opts.preload.GetPreload())
	}
	return tx
}

func NewBuilder() Builder {
	return &builder{}
}
