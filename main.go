package main

import "fmt"

func main() {
	blog := New()

	fmt.Println(blog)

	blog.SaveArticle(Article{
		"My first blog post",
		"Today, we will be talking about blogging",
	})

	fmt.Println(blog)
}
