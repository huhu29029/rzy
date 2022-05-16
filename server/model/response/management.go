package response

import (
	"online-exam/model"
	"online-exam/model/request"
)

type GetQuestionsResponse struct {
	Total     uint              `json:"total"`
	Questions []model.Questions `json:"questions"`
}

type AddQuestionResponse struct {
	Question model.Questions `json:"question"`
}

type DelQuestionResponse struct {
	Question struct{} `json:"question"`
}
type DelQuestionFromQuestionBook struct {
	WrongQuestionBook struct{} `json:"question"`
}
type ExamResponse struct {
	Exam request.Exam `json:"exam"`
}

type SubmitExamResponse struct {
	ID uint `json:"id"`
}
