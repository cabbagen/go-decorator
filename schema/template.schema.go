package schema

type TemplateSchema struct {
	BaseSchema
	Name            string     `gorm:"column:name; type:varchar(255); not null default: \"\"" json:"name"`
	Cover           string     `gorm:"column:cover; type:varchar(255); not null default: \"\"" json:"cover"`
	ProjectId       int        `gorm:"column:project_id; type:int; not null" json:"projectId"`
	CategoryId      int        `gorm:"column:category_id; type:int; not null" json:"categoryId"`
}
