package response

import (
	"online-exam/model"
	"time"
)

type GetExamHistoryResponse struct {
	Total       uint                `json:"total"`
	ExamHistory []model.ExamHistory `json:"examHistory"`
}

type GetExamsResponse struct {
	Total uint          `json:"total"`
	Exams []model.Exams `json:"exams"`
}

type ExamResult struct {
	ID         uint                   `json:"id"`
	Title      string                 `json:"title"`
	Score      uint                   `json:"score"`
	TotalScore uint                   `json:"totalScore"`
	StartTime  time.Time              `json:"startTime"`
	EndTime    time.Time              `json:"endTime"`
	Questions  []model.QuestionResult `json:"questions"`
}

type ExamRecommendations struct {
	QuestionIds []int `json:"questionIds"`
}
type ExamResultResponse struct {
	Result          ExamResult          `json:"result"`
	Recommendations ExamRecommendations `json:"recommendations"`
}
