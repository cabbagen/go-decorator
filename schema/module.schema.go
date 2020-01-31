package schema

type ModuleSchema struct {
	BaseSchema
	PageId               int           `gorm:"column:page_id; type:int; not null" json:"pageId"`
	SortNo               int           `gorm:"column:sort_no; type:int; not null" json:"sortNo"`
	Content              string        `gorm:"column:content; type:text" json:"content"`
}

type ModuleSort struct {
	Id                   int           `json:"id"`
	SortNo               int           `json:"sortNo"`
}