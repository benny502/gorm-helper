package helper

import (
	"github.com/benny502/gorm-helper/associate"
	"gorm.io/gorm"
)

type Builder interface {
	WithWhere(where Where) Builder
	WithAssociate(associate associate.Associate) Builder
	WithPreload(preload ...associate.Preload) Builder
	WithHaving(having ...Having) Builder
	WithSelect(field ...Select) Builder
	WithGroup(group Group) Builder
	Build(db *gorm.DB) *gorm.DB
}

type options struct {
	where     Where
	associate associate.Associate
	preloads  []associate.Preload
	group     Group
	having    []Having
	field     []Select
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

func (b builder) WithPreload(preloads ...associate.Preload) Builder {
	b.opts.preloads = preloads
	return &b
}

func (b builder) WithGroup(group Group) Builder {
	b.opts.group = group
	return &b
}

func (b builder) WithHaving(having ...Having) Builder {
	b.opts.having = having
	return &b
}

func (b builder) WithSelect(field ...Select) Builder {
	b.opts.field = field
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
		b.opts.where.Rewind()
	}

	if b.opts.associate != nil {
		tx.Preload(b.opts.associate.GetPreload())
		joins := b.opts.associate.GetJoinsString()
		for _, join := range joins {
			tx.Joins(join)
		}
	}

	if len(b.opts.preloads) > 0 {
		for _, preload := range b.opts.preloads {
			tx.Preload(preload.GetPreload(), preload.GetArgs()...)
		}
	}

	if b.opts.group != nil {
		tx.Group(b.opts.group.GetGroup())
	}

	if len(b.opts.having) > 0 {
		for _, having := range b.opts.having {
			tx.Having(having.GetQuery(), having.GetArgs()...)
		}
	}

	if len(b.opts.field) > 0 {
		for _, field := range b.opts.field {
			tx.Select(field.GetQuery(), field.GetArgs()...)
		}
	}
	return tx
}

func NewBuilder() Builder {
	return &builder{}
}
