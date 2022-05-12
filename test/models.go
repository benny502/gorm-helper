package test

import "time"

type User struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Username       string    `gorm:"column:username;type:varchar(255)" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Email          string    `gorm:"column:email;type:varchar(255)" json:"email"`
	CreatedAt      time.Time `gorm:"column:createdAt;type:datetime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;type:datetime" json:"updatedAt"`
	RealLastName   string    `gorm:"column:realLastName;type:varchar(255)" json:"realLastName"`
	Phone          string    `gorm:"column:phone;type:varchar(100)" json:"phone"`
	Status         int       `gorm:"column:status;type:tinyint(4);default:0" json:"status"` // 0正常 1关闭 2 未验证
	RealFirstName  string    `gorm:"column:realFirstName;type:varchar(255)" json:"realFirstName"`
	Nickname       string    `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
	Usercode       string    `gorm:"column:usercode;type:varchar(255)" json:"usercode"`
	IsAllCamp      int       `gorm:"column:isAllCamp;type:tinyint(1);default:0" json:"isAllCamp"` // 1为全部  默认0
	LastLoginTime  time.Time `gorm:"column:lastLoginTime;type:datetime" json:"lastLoginTime"`
	LoginStatus    int       `gorm:"column:loginStatus;type:tinyint(4);default:0" json:"loginStatus"` // 0 未登录 1 登录
	LastActionTime time.Time `gorm:"column:lastActionTime;type:datetime" json:"lastActionTime"`
	Type           int       `gorm:"column:type;type:tinyint(4);default:11" json:"type"` // 1付费用户（团队）、2付费附属、3付费用户（个人）、4试用用户（团队）、5试用附属、6试用用户（个人）、7免费用户（HY团队）、8免费附属(HY附属)、9免费用户（HY个人）、10免费用户（RH）、11潜在用户、12增值用户
	DeactivateTime time.Time `gorm:"column:deactivateTime;type:datetime" json:"deactivateTime"`
	Industry       int       `gorm:"column:industry;type:int(11)" json:"industry"`
	Job            int       `gorm:"column:job;type:int(11)" json:"job"`
	Organize       string    `gorm:"column:organize;type:varchar(255)" json:"organize"`
	Remark         string    `gorm:"column:remark;type:varchar(255)" json:"remark"` // 备注
	PrevLoginTime  time.Time `gorm:"column:prevLoginTime;type:datetime" json:"prevLoginTime"`
	PermissionType int       `gorm:"column:permissionType;type:tinyint(4);default:0" json:"permissionType"` // 0 默认权限 1 附属权限 2 分配权限
	UserStatus     int       `gorm:"column:userStatus;type:tinyint(4);default:0" json:"userStatus"`         // 0 正常  1新客  2 活跃 3回归 4 僵尸 5 待激活 6 注销 7系统注销
	BackTime       time.Time `gorm:"column:backTime;type:datetime" json:"backTime"`
	Version        int       `gorm:"column:version;type:tinyint(4);default:0" json:"version"`           // 0基础版 1高级版
	IsTrial        int       `gorm:"column:isTrial;type:tinyint(4);default:0;NOT NULL" json:"isTrial"`  // 1 使用 0 正常
	RequestTime    time.Time `gorm:"column:requestTime;type:datetime" json:"requestTime"`               // 申请时间
	ReviewStatus   int       `gorm:"column:reviewStatus;type:tinyint(4);default:0" json:"reviewStatus"` // 1申请试用中
	TeamUser       TeamUser
}

func (m *User) TableName() string {
	return "m_user"
}

type Team struct {
	Id                  int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name                string    `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt           time.Time `gorm:"column:createdAt;type:datetime" json:"createdAt"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;type:datetime" json:"updatedAt"`
	TeamCode            string    `gorm:"column:teamCode;type:varchar(255)" json:"teamCode"`
	ExpireTime          time.Time `gorm:"column:expireTime;type:datetime" json:"expireTime"`
	Type                int       `gorm:"column:type;type:tinyint(4);default:0" json:"type"`           // 0 基础版 1 高级版 2尊享版  3HY版 4 HY(rh)
	IsTrial             int       `gorm:"column:isTrial;type:tinyint(4);default:0" json:"isTrial"`     // 0 否 1 是
	ValidityPeriod      int       `gorm:"column:validityPeriod;type:tinyint(4)" json:"validityPeriod"` // 1 14天 2 一个月 3三个月 4 一年 5 两年 6三年 7四年 8五年 9自定义
	Capacity            int       `gorm:"column:capacity;type:tinyint(4)" json:"capacity"`             // 容量 1 1G 2 2G 3 10G 4 20G 5 50G 6 100G 7 150G 8 200G 9 300G
	IsRH                int       `gorm:"column:isRH;type:tinyint(4);default:0;NOT NULL" json:"isRH"`
	ValidityPeriodStart time.Time `gorm:"column:validityPeriodStart;type:datetime" json:"validityPeriodStart"`
	ValidityPeriodEnd   time.Time `gorm:"column:validityPeriodEnd;type:datetime" json:"validityPeriodEnd"`
	Discount            int       `gorm:"column:discount;type:tinyint(4)" json:"discount"`                  // 折扣
	DiscountValue       int       `gorm:"column:discountValue;type:int(11)" json:"discountValue"`           // 折扣差额
	VersionLevelTotal   int       `gorm:"column:versionLevelTotal;type:int(11)" json:"versionLevelTotal"`   // 版本小计
	CapacityAdditional  int       `gorm:"column:capacityAdditional;type:int(11)" json:"capacityAdditional"` // 附加容量
	CapacityTotal       int       `gorm:"column:capacityTotal;type:int(11)" json:"capacityTotal"`           // 容量小计
	Seats               int       `gorm:"column:seats;type:tinyint(4)" json:"seats"`                        // 席位 1 1人 2 10人 3 20人 4 50人 5 100人 6 150人 7 200人 8
	SeatsAdditional     int       `gorm:"column:seatsAdditional;type:int(11)" json:"seatsAdditional"`       // 附加席位
	SeatsTotal          int       `gorm:"column:seatsTotal;type:int(11)" json:"seatsTotal"`                 // 席位小计
	Permission          string    `gorm:"column:permission;type:text" json:"permission"`                    // 权限
	CapacityValue       int       `gorm:"column:capacityValue;type:int(11)" json:"capacityValue"`           // 总容量
	SeatsValue          int       `gorm:"column:seatsValue;type:int(11)" json:"seatsValue"`                 // 总座位数
	HYType              int       `gorm:"column:HYType;type:tinyint(4);default:1" json:"HYType"`            // 1 正式（免费） 2 正式（RH）
	RHHidden            string    `gorm:"column:RHHidden;type:varchar(255)" json:"RHHidden"`
	DistributionAt      time.Time `gorm:"column:distributionAt;type:datetime" json:"distributionAt"`
	IsHide              int       `gorm:"column:isHide;type:tinyint(4)" json:"isHide"`                                 // RH是否隐藏信息0 否 1是
	PermissionOn        int       `gorm:"column:permissionOn;type:tinyint(4);default:0" json:"permissionOn"`           // 0 关闭权限 1 开启权限
	PermissionChanged   int       `gorm:"column:permissionChanged;type:tinyint(4);default:0" json:"permissionChanged"` // 1 权限被修改
}

func (m *Team) TableName() string {
	return "m_team"
}

type TeamUser struct {
	Id        int64     `gorm:"column:id;type:bigint(20);" json:"id"`
	UserId    int64     `gorm:"column:userId;type:bigint(20);primary_key;NOT NULL" json:"userId"`
	TeamId    int64     `gorm:"column:teamId;type:bigint(20);primary_key;NOT NULL" json:"teamId"`
	Role      int       `gorm:"column:role;type:tinyint(4);default:1" json:"role"` // 1 成员  2团队管理员 3团队所有者
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:datetime" json:"updatedAt"`
	Status    int       `gorm:"column:status;type:tinyint(4)" json:"status"` // 0 审批中  1申请通过  2申请失败 3 已撤销
	Team      Team      `gorm:"foreignKey:TeamId;references:Id"`
}

func (m *TeamUser) TableName() string {
	return "m_team_user"
}
