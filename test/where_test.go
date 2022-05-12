package test

import (
	"fmt"
	"testing"

	"github.com/benny502/gorm-helper/where"
	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	where := where.NewWhere()
	where.Add("m_user.id in ?", []int{1, 2, 3})
	where.Add("m_user.userName like ?", "%admin%")
	for where.Next() {
		sql := where.GetQuery()
		assert.NotEmpty(t, sql)
		args := where.GetArgs()
		assert.NotEmpty(t, args)
		fmt.Printf("%s\n%v", sql, args)
	}
	//panic("doom")
}
