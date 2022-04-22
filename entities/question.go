package entities

type Question struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}
