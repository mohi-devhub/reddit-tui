package models

type Post struct {
	Title     string `json:"title"`
	Subreddit string `json:"subreddit"`
	Author    string `json:"author"`
	Upvotes   int    `json:"upvotes"`
	Comments  int    `json:"comments"`
}
