package schema

type TemplateCategorySchema struct {
	BaseSchema
	Name            string     `gorm:"column:name; type:varchar(255); not null default: \"\"" json:"name"`
}
