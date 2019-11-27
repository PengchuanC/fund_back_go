package models

type News struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	Abstract string `gorm:"column:abstract" json:"abstract"`
	Url      string `gorm:"column:url" json:"url"`
	Source   string `gorm:"column:source" json:"source"`
	Keyword  string `gorm:"column:keyword" json:"keyword"`
	SaveDate string `gorm:"column:savedate;type:datetime" json:"savedate"`
}

func (n *News) AfterFind() {
	n.SaveDate = n.SaveDate[:10]
}
