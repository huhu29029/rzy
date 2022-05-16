package request

import (
	"online-exam/global"
	"online-exam/model"
	"time"
)

type GetRandomExamResponse struct {
	Exam model.Exams `json:"exam"`
}

type Exam struct {
	global.MODEL
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	global.MODEL
	Title   string           `json:"title"`
	Type    uint             `json:"type"`
	Options []QuestionOption `json:"options"`
}

type QuestionOption struct {
	global.MODEL
	Title string `json:"title"`
}

type ExamAnswer struct {
	ExamId    uint                 `json:"examId" gorm:"comment:试卷ID"`
	StartTime time.Time            `json:"startTime" gorm:"comment:答题开始时间戳"`
	EndTime   time.Time            `json:"endTime" gorm:"comment:答题结束时间戳"`
	Questions []ExamAnswerQuestion `json:"questions"`
}

type ExamAnswerQuestion struct {
	ID      uint   `json:"id"`
	Options []uint `json:"options"`
}
