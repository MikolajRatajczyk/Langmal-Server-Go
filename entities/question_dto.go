package entities

type QuestionDto struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}
