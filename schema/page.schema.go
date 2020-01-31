package schema

type PageSchema struct {
	BaseSchema
	ProjectId          int       `gorm:"column:project_id; type:int; not null" json:"projectId"`
	Name               string    `gorm:"column:name; type:varchar(255); not null default: 新建页面" json:"name"`
	Link               string    `gorm:"column:link; type:varchar(255); not null default: \"\"" json:"link"`
	CoverImg           string    `gorm:"column:cover_img; type:varchar(255); not null default: \"\"" json:"coverImg"`
}