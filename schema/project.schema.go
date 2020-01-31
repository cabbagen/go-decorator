package schema

type ProjectSchema struct {
	BaseSchema
	Name            string     `gorm:"column:name; type:varchar(255); not null default: \"\"" json:"name"`
	Type            int        `gorm:"column:type; type:samllint; not null default: 1" json:"type"`
	State           int        `gorm:"column:state; type:samllint; not null default: 1" json:"state"`
}
