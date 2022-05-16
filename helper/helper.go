package helper

import (
	"github.com/benny502/gorm-helper/associate"
	"github.com/benny502/gorm-helper/builder"
	"github.com/benny502/gorm-helper/where"
	"gorm.io/gorm/schema"
)

func NewAssociate(model schema.Tabler, preload string) builder.Associate {
	return associate.NewAssociate(model, preload)
}

func NewBuilder() builder.Builder {
	return builder.NewBuilder()
}

func NewWhere() builder.Where {
	return where.NewWhere()
}
