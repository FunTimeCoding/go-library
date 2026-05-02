package argument

type CreateView struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Syntax  string `json:"syntax"`
}
