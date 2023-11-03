package models

type QuestionEntity struct {
	Title   string
	Id      string `gorm:"primaryKey"`
	Options StorableStringArray
	Answer  string
	// foreign key, must match QuizEntity.Id
	QuizEntityId string
}
