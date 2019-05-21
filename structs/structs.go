package structs

type CodeSnippet struct {
	ID          int 	`json:"id"`
	Title       string	`json:"title"`
	Codelines   string	`json:"codelines"`
	Tag         string	`json:"tag"`
	Author      string	`json:"author"`
	ContentType string	`json:"contenttype"`
	CreatedAt   string	`json:"createdat"`
	UpdatedAt   string	`json:"updatedat"`
}

type CodeSnippetArray struct {
	CodeSnippetArray []CodeSnippet	`json:"codesnippet"`
}
