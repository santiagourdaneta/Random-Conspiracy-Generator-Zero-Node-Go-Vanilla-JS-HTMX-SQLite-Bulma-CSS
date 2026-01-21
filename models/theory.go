package models

type Theory struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"` // URL amigable: la-nasa-oculta-queso
}
