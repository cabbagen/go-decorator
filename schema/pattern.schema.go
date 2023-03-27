package schema

type PatternSchema struct {
	BaseSchema
	Name               string    `gorm:"column:name; type:varchar(255); not null" json:"name"`
	Code               string    `gorm:"column:code; type:varchar(255); not null; unique" json:"code"`
}
