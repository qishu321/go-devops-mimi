package nav

// 代表单个导航链接
type Link struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"` // 链接名称
	Desc string `gorm:"size:255" json:"desc"`          // 链接描述
	URL  string `gorm:"size:255;not null" json:"url"`  // 跳转地址
}

// 代表一个导航分类，例如：监控、CI/CD
type Nav struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `gorm:"size:50;unique;not null" json:"name"` // 分类名称
	NavSort int     `gorm:"default:0" json:"nav_sort"`           // 同分类下排序
	Links   []*Link `gorm:"many2many:t_link_s;comment:'关联的导航链接'" json:"t_link_s"`
}

func (m *Link) TableName() string {
	return "t_link"
}

func (m *Nav) TableName() string {
	return "t_nav"
}
