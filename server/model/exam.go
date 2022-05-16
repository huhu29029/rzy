package model

import (
	"online-exam/global"
	"time"
)

// 试卷
type Exams struct {
	global.MODEL
	Title     string      `json:"title" gorm:"comment:考试名称"`
	UserId    uint        `json:"userId" gorm:"comment:用户ID"`
	Questions []Questions `json:"questions" gorm:"many2many:exams_questions;"`
	SysUser   SysUser     `json:"-" gorm:"foreignKey:UserId;references:ID"`
}

// 用户答卷基本信息
type ExamAnswers struct {
	global.MODEL
	StartTime time.Time             `json:"startTime" gorm:"comment:答题开始时间戳"`
	EndTime   time.Time             `json:"endTime" gorm:"comment:答题结束时间戳"`
	UserId    uint                  `json:"userId" gorm:"comment:用户ID"`
	ExamId    uint                  `json:"examId" gorm:"comment:试卷ID"`
	Questions []ExamAnswerQuestions `json:"questions" gorm:"foreignKey:ExamAnswerId;references:ID"`
	Exams     Exams                 `json:"-" gorm:"foreignKey:ExamId;references:ID;comemnt:对应试卷"`
	SysUser   SysUser               `json:"-" gorm:"foreignKey:UserId;references:ID;comment:对应用户"`
}

// 用户答卷题目存储
type ExamAnswerQuestions struct {
	global.MODEL
	ExamAnswerId    uint                        `json:"-" gorm:"comemnt:用户答卷ID"`
	QuestionId      uint                        `json:"questionId" gorm:"comment:题目ID"`
	QuestionOptions []ExamAnswerQuestionOptions `json:"options" gorm:"foreignKey:ExamAnswerQuestionId;references:ID"`
	Questions       Questions                   `json:"-" gorm:"foreignKey:QuestionId;references:ID;comemnt:对应问题"`
}

// 用户答卷个题目选项
type ExamAnswerQuestionOptions struct {
	global.MODEL
	ExamAnswerQuestionId uint            `json:"-" gorm:"comment:用户答卷题目"`
	QuestionOptionId     uint            `json:"questionOptionId" gorm:"comment:问题选项ID"`
	QuestionOptions      QuestionOptions `json:"-" gorm:"foreignKey:QuestionOptionId;references:ID;comment:对应问题选项"`
}

type ExamHistory struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	ExamId    uint      `json:"examId"`
	ExamTitle string    `json:"examTitle"`
}

type ExamsQuestions struct {
	ExamsId     uint `json:"examsId"`
	QuestionsId uint `json:"questionsId"`
}
