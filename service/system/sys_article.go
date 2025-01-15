package system

import (
	"blog-backend/global"
	"blog-backend/model/system"
)

type ArticleService struct{}

func (a *ArticleService) CreateArticle(art *system.Article) (err error) {
	if err = global.YAGAMI_DB.Create(art).Error; err != nil {
		return err
	}
	return
}

func (a *ArticleService) DeleteArticleById() {

}

func (a *ArticleService) UpdateArticleById() {

}

func (a *ArticleService) GetArticleList() {

}

func (a *ArticleService) GetArticleByTitle() {

}

func (a *ArticleService) GetArticleById() {

}
