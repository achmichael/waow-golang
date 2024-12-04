package dtos

type ArticleRequest struct {
	Category_id string `json:"category_id"`
	User_id     string `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Slug        string `json:"slug"`
	View_count  int    `json:"view_count"`
}
