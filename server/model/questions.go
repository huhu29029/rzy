package model

import (
	"online-exam/global"
	"time"
)

// 题目对应选项库
type QuestionOptions struct {
	global.MODEL
	Title      string `json:"title" gorm:"comment:选项题干"`
	IsRight    bool   `json:"isRight" gorm:"comment:是否正确"`
	QuestionId uint   `json:"-" gorm:"comment:对应题目"`
}

//题目库
type Questions struct {
	global.MODEL
	Title     string            `json:"title" form:"title" gorm:"comment:题干"`
	Analysis  string            `json:"analysis" form:"analysis" gorm:"comment:详解"`
	Type      uint              `json:"type" gorm:"comment:题目类型，单选1 判断2 多选3"`
	ChapterId uint              `json:"chapterId" gorm:"comment:所属章节"`
	UserId    uint              `json:"userId" gorm:"comment:用户ID"`
	Options   []QuestionOptions `json:"options" gorm:"foreignKey:QuestionId;references:ID;comment:题目选项"`
	Tags      []Tags            `json:"tags" gorm:"many2many:questions_tags"`
	Chapters  Chapters          `json:"-" gorm:"foreignKey:ChapterId;references:ID"`
	SysUser   SysUser           `json:"-" gorm:"foreignKey:UserId;references:ID"`
}

//错题库
type WrongQuestionBook struct {
	global.MODEL
	UserId     uint   `json:"userId" gorm:"comment:用户ID"`
	QuestionID uint   `json:"id"`
	Title      string `json:"title" form:"title" gorm:"comment:题干"`
	Analysis   string `json:"analysis" form:"analysis" gorm:"comment:详解"`
	Type       uint   `json:"type" gorm:"comment:题目类型，单选1 判断2 多选3"`
	IsRight    bool   `json:"isRight" comment:"用户回答是否正确"`
}

// 用户答题库
type Answers struct {
	ID         uint            `gorm:"primaryKey"`
	StartTime  time.Time       `json:"startTime" gorm:"comment:答题开始时间戳"`
	EndTime    time.Time       `json:"endTime" gorm:"comment:答题结束时间戳"`
	UserId     uint            `json:"userId" gorm:"foreignKey:id;comment:用户ID"`
	QuestionId uint            `json:"questionId" gorm:"comment:题目ID"`
	IsRight    bool            `json:"isRight" gorm:"comment:答题正误"`
	Questions  Questions       `json:"-" gorm:"foreignKey:QuestionId;references:ID;comment:对应问题"`
	SysUser    SysUser         `json:"-" gorm:"foreignKey:UserId;references:ID;comment:对应用户"`
	Options    []AnswerOptions `json:"options" gorm:"foreignKey:AnswerId;references:ID;comment:对应用户答题"`
}

// 用户答题对应选项库
type AnswerOptions struct {
	ID               uint            `gorm:"primaryKey"`
	AnswerId         uint            `json:"answerId" gorm:"comment:问题ID"`
	QuestionOptionId uint            `json:"questionOptionId" gorm:"comment:问题选项ID"`
	QuestionOptions  QuestionOptions `json:"questionOptions" gorm:"foreignKey:QuestionOptionId;references:ID;comment:题目选项"`
}

// 标签库
type Tags struct {
	global.MODEL
	Layer1 string `json:"layer1"`
	Layer2 string `json:"layer2"`
	Layer3 string `json:"layer3"`
	Layer4 string `json:"layer4"`
}

// 标签-问题 many2many
type QuestionsTags struct {
	QuestionsId uint
	TagsId      uint
}

// 章节库
type Chapters struct {
	global.MODEL
	Title string `json:"title"`
}

// 答题历史
type QuestionHistory struct {
	ID             uint           `json:"id"`
	StartTime      time.Time      `json:"startTime"`
	EndTime        time.Time      `json:"endTime"`
	QuestionId     uint           `json:"questionId"`
	QuestionResult QuestionResult `json:"questionResult"`
}

// 一次答题结果
type QuestionResult struct {
	Title    string                 `json:"title"`
	Analysis string                 `json:"analysis"`
	Type     uint                   `json:"type"`
	IsRight  bool                   `json:"isRight" comment:"用户回答是否正确"`
	Options  []QuestionOptionResult `json:"options"`
}

type QuestionOptionResult struct {
	Title     string `json:"title"`
	IsRight   bool   `json:"isRight"`
	IsChoosed bool   `json:"isChoosed"`
}
