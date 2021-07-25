package schema

type ModuleSchema struct {
	BaseSchema
	PageId               int           `gorm:"column:page_id; type:int; not null" json:"pageId"`
	Type                 string        `gorm:"column:type; type:varchar; not null; default:\"\"" json:"type"`
	SortNo               int           `gorm:"column:sort_no; type:int; not null" json:"sortNo"`
	Content              string        `gorm:"column:content; type:text" json:"content"`
}

type ModuleSort struct {
	Id                   int           `json:"id"`
	SortNo               int           `json:"sortNo"`
}

type PageModule struct {
	ModuleSchema
	PageName             string        `gorm:"column:page_name; type:varchar" json:"pageName"`
	PageLink             string        `gorm:"column:page_link; type:varchar" json:"pageLink"`
}
