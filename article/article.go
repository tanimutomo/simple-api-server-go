package article

type Article struct {
	Title       string `json:"title"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type ArticleList struct {
	Articles []Article
}

type Articles interface {
	Add(Article)
	GetAll() []Article
}

func (r *ArticleList) Add(article Article) {
	r.Articles = append(r.Articles, article)
}

func (r *ArticleList) GetAll() []Article {
	return r.Articles
}
