package glue

import (
	"onsite/model"
)

type ArticleStoreInterface interface {
	GetBySlug(string) (*model.Article, error)
	GetUserArticleBySlug(userID uint, slug string) (*model.Article, error)
	CreateArticle(*model.Article) error
	UpdateArticle(*model.Article, []string) error
	DeleteArticle(*model.Article) error
	List(offset, limit int) ([]model.Article, int, error)
	ListByTag(tag string, offset, limit int) ([]model.Article, int, error)
	ListByAuthor(username string, offset, limit int) ([]model.Article, int, error)
	ListByWhoFavorited(username string, offset, limit int) ([]model.Article, int, error)
	ListFeed(userID uint, offset, limit int) ([]model.Article, int, error)
	AddComment(*model.Article, *model.Comment) error
	GetCommentsBySlug(string) ([]model.Comment, error)
	GetCommentByID(uint) (*model.Comment, error)
	DeleteComment(*model.Comment) error
	AddFavorite(*model.Article, uint) error
	RemoveFavorite(*model.Article, uint) error
	ListTags() ([]model.Tag, error)
}

type UserStoreInterface interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	AddFollower(user *model.User, followerID uint) error
	RemoveFollower(user *model.User, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}
