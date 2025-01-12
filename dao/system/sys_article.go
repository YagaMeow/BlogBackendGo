package system

type Article struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	UpdateTime string `json:"update_time"`
	CreateTime string `json:"create_time"`
	Abstract   string `json:"abstract"`
	Content    string `json:"content"`
}

func (a *Article) GetId() int {
	return a.Id
}

func (a *Article) GetTitle() string {
	return a.Title
}

func (a *Article) GetContent() string {
	return a.Content
}

func (a *Article) GetAbstract() string {
	return a.Abstract
}

func (a *Article) GetCreateTime() string {
	return a.CreateTime
}

func (a *Article) TableName() string {
	return "sys_article"
}
