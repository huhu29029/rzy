package request

import (
	"time"
)

type Answer struct {
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	QuestionId uint      `json:"questionId"`
	Options    []uint    `json:"options"`
}

type GetQuestionByChaptersRequest struct {
	ChapterIds []uint `json:"chapterIds"`
}
