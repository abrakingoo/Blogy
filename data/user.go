package data

import (
	"time"
)

type Blog struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Blog_content string `json:"blog_content"`
}

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Create_at time.Time `json:"created_at"`
	Blogs     []Blog    `json:"blogs"`
}

// Source represents the source of an article.
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description *string `json:"description"` // Nullable field
	URL         string `json:"url"`
	Image    *string `json:"image"` // Nullable field
	PublishedAt string `json:"publishedAt"`
	Content     *string `json:"content"` // Nullable field
}

// Article represents a news article with various attributes.
type Article struct {
	Status string `json:"status"`
	Totalresults int `json:"totalresults"`
	Articles []Source
}
