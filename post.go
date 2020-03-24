package main

import "github.com/google/uuid"

// Post Model
type Post struct {
	id      uuid.UUID
	content string
}

var posts []Post

func (Post) getTableData() []Post {
	return posts
}

// This method MUST be implemented by an object
// to be qualified as a [Model]
func (post Post) getTableName() string {
	return "posts"
}

func (post Post) initialiseTable() {
	posts = []Post{
		Post{
			id:      uuid.New(),
			content: "sample post 1",
		},
		Post{
			id:      uuid.New(),
			content: "sample post 2",
		},
	}
}
