package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type QuizRepoInterface interface {
	Create(quiz models.QuizEntity) bool
	FindAll() []models.QuizEntity
	Find(id string) (models.QuizEntity, bool)
}

func NewQuizRepo(dbName string) QuizRepoInterface {
	quizRepo := &quizRepo{
		db: getDb(dbName, models.QuizEntity{}, models.QuestionEntity{}),
	}

	noQuizzes := len(quizRepo.FindAll()) == 0
	if noQuizzes {
		quizRepo.Create(models.DefaultQuiz1())
		quizRepo.Create(models.DefaultQuiz2())
	}

	return quizRepo
}

type quizRepo struct {
	db *gorm.DB
}

func (qr *quizRepo) Create(quiz models.QuizEntity) bool {
	if err := qr.db.Create(&quiz).Error; err != nil {
		log.Println("Failed to create a new quiz in DB!")
		return false
	} else {
		return true
	}
}

func (qr *quizRepo) FindAll() []models.QuizEntity {
	var quizzes []models.QuizEntity

	err := qr.db.
		Preload("Questions").
		Find(&quizzes).
		Error
	if err != nil {
		return []models.QuizEntity{}
	}

	return quizzes
}

func (qr *quizRepo) Find(id string) (models.QuizEntity, bool) {
	var quiz models.QuizEntity
	err := qr.db.
		Preload("Questions").
		First(&quiz, "id = ?", id).
		Error
	success := err == nil
	return quiz, success
}
