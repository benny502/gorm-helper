# gorm-helper
A helper for go-gorm


当我们在Gorm中要关联两张表时我们可以使用Preload

```Go
db.Preload("TeamUser").Where("userId in ?", []int{1, 2, 3}).Find(&user)

//SELECT * FROM `m_user` WHERE id in (1,2,3)

//SELECT * FROM `m_team_user` WHERE `m_team_user`.`userId` IN (1,2,3)

```

TeamUser的结果会自动填充到User结构体的TeamUser属性中

有时候我们需要根据关联表联合设置查询条件，Gorm提供了Joins

```Go

db.Joins("TeamUser").Where("m_user.id in ? and m_team_user.status = 1", []int{1, 2, 3}).Find(&user)

// SELECT `m_user`.`id`,`m_user`.`username`,`m_user`.`password`,`m_user`.`email`,`m_user`.`createdAt`,`m_user`.`updatedAt`,`TeamUser`.`id` AS `TeamUser__id`,`TeamUser`.`userId` AS `TeamUser__userId`,`TeamUser`.`teamId` AS `TeamUser__teamId`,`TeamUser`.`role` AS `TeamUser__role`,`TeamUser`.`createdAt` AS `TeamUser__createdAt`,`TeamUser`.`updatedAt` AS `TeamUser__updatedAt`,`TeamUser`.`status` AS `TeamUser__status` FROM `m_user` 
// LEFT JOIN `m_team_user` `TeamUser` ON `m_user`.`id` = `TeamUser`.`userId` WHERE m_user.id in (1,2,3) and m_team_user.status = 1

```Go

数据也会自动填充User.TeamUser, 很完美

但如果我们再加一张关联表呢，Preload支持嵌套查询，但是我们无法用关联表的条件来进行筛选

```Go

db.Preload(clause.Associations).Preload("TeamUser.Team").Where("TeamUser.row = ?", 3).Find(&user)

// SELECT * FROM `m_user` WHERE TeamUser.row = 3

// Error 1054: Unknown column 'TeamUser.row' in 'where clause'

```

如果我们使用Joins，我们必须自定义一个临时结构体，有没有什么方法能两者兼顾呢？

```Go

db.Preload("TeamUser.Team").
		Joins("LEFT JOIN m_team_user ON m_user.id=m_team_user.userId").
		Joins("LEFT JOIN m_team ON m_team_user.teamId=m_team.id").
		Where("m_user.id in ?", []int{1, 2, 3}).Find(&user)
    
// SELECT * FROM `m_team` WHERE `m_team`.`id` IN (1,2)

// SELECT * FROM `m_team_user` WHERE `m_team_user`.`userId` IN (2,3)

// SELECT `m_user`.`id`,`m_user`.`username`,`m_user`.`password`,`m_user`.`email`,`m_user`.`createdAt`,`m_user`.`updatedAt`
// FROM `m_user` 
// LEFT JOIN m_team_user ON m_user.id=m_team_user.userId 
// LEFT JOIN m_team ON m_team_user.teamId=m_team.id WHERE m_user.id in (1,2,3) and m_team_user.role=3

```

上面这种方式可以完美解决，代价是多了两条sql语句，我们不用再自定义结构体了，Gorm会帮我们将结果自动填充进User.TeamUser与User.TeamUser.Team里

可是手动构造Join语句还是太累了，于是我对上面的方式进行了封装

```Go

import (
    "github.com/benny502/gorm-helper/builder"
    "github.com/benny502/gorm-helper/associate"
)

builder.NewBuilder().
  WithAssociate(associate.NewAssociate(&test.User{}, "TeamUser.Team")).
  Build(db).
  Where("m_user.id in ?", []int{1, 2, 3}).
  Find(&user)
  
```

这条语句等价于上面的，Builder会自动分析TeamUser与Team的实体类，自动生成Join语句









