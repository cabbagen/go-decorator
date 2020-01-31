package schema

type UserSchema struct {
	BaseSchema
	Username             string     `gorm:"column:username; type:varchar(255); not null; unique key" json:"username"`
	Nickname             string     `gorm:"column:nickname; type:varchar(255); not null" json:"nickname"`
	Password             string     `gorm:"column:password; type:varchar(255); not null" json:"password"`
	Avatar               string     `gorm:"column:avatar; type:varchar(255); not null" json:"avatar"`
	Mobile               string     `gorm:"column:mobile; type:varchar(255); not null; unique key" json:"mobile"`
	Email                string     `gorm:"column:email; type:varchar(255); not null; unique key" json:"email"`
}
