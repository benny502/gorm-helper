package main

import (
	"github.com/benny502/gorm-helper/associate"
	"github.com/benny502/gorm-helper/builder"
	"github.com/benny502/gorm-helper/test"
	"github.com/benny502/gorm-helper/where"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func main() {
	dsn := "root:root@tcp(192.168.81.13:3306)/hy_flexbi_online?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	where := where.NewWhere()
	where.Add("m_user.id in ?", []int{1, 2, 3})

	user := []*test.User{}
	db.Model(&test.User{})
	builder.NewBuilder().WithWhere(where).WithAssociate(associate.NewAssociate(&test.User{}, "TeamUser.Team")).Build(db).Find(&user)

}
