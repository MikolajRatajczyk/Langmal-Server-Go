package models

type QuestionDto struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"`
}
