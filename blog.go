package main

// Article is a struct type that defines an article in the blog.
type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Blog represents the blog itself and consists of an array of `Article`s.
type Blog struct {
	Articles []Article
}

// New is a funciton to instantiate a new `Blog` and returns a
// pointer to a Blog type.
func New() *Blog {
	return &Blog{}
}

// SaveArticle is a method to add a new article to a blog's article collection.
// This is a method on the `Blog` type. Note the receiver of type
// `Blog` with the name `b`.
func (b *Blog) SaveArticle(article Article) {
	b.Articles = append(b.Articles, article)
}

// FetchAll is a method to retrieve all articles in a `Blog` instance.
// Returns an array of `Articles`.
// This is another method on the `Blog` type.
func (b *Blog) FetchAll() []Article {
	return b.Articles
}
