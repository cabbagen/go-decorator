package schema

import "time"

type BaseSchema struct {
	ID               int            `gorm:"column:id; type:int; primary key; not null" json:"id"`
	Remark           string         `gorm:"column:remark; type:varchar(255); not null default: \"\"" json:"remark"`
	CreatedAt        time.Time      `gorm:"column:created_at; type:datetime; default now()" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at; type:datetime; default now()" json:"updatedAt"`
}
