package models

type ResultEntity struct {
	Correct   int
	Wrong     int
	QuizId    string
	CreatedAt int64
	UserId    string
}
