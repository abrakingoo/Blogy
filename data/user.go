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
