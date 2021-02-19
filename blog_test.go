package main

import "testing"

func TestSaveArticle(t *testing.T) {
	blog := New()

	blog.SaveArticle(Article{"Title 1", "Title 1's Body"})

	if blog.Articles[0].Title != "Title 1" {
		t.Errorf("Article was not added")
	}
}

func TestFetchAllArticles(t *testing.T) {
	blog := New()

	blog.SaveArticle(Article{"Title 1", "Title 1's Body"})

	articles := blog.FetchAll()

	if len(articles) == 0 {
		t.Errorf("Error fetching all articles")
	}
}
