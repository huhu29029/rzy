package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"online-exam/core"
	"online-exam/global"
	"online-exam/initialize"
	"online-exam/model"
	"online-exam/model/response"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func GetExamResult(id string) ([2]uint, error) {
	// 获取用户答卷
	examAnswerId, _ := strconv.Atoi(id)

	var examAnswer model.ExamAnswers

	var examAnswerRes *gorm.DB
	examAnswerRes = global.DB.Where(&model.ExamAnswers{
		MODEL: global.MODEL{
			ID: uint(examAnswerId),
		},
	}).First(&examAnswer)

	if examAnswerRes.Error != nil {
		log.Println(id, examAnswerRes.Error)
		return [2]uint{}, examAnswerRes.Error
	}

	// 获取试卷信息
	var exam model.Exams

	result := global.DB.First(&exam, examAnswer.ExamId)

	if result.Error != nil {
		log.Println(id, examAnswer.ExamId, result.Error)
		return [2]uint{}, result.Error
	}

	global.DB.Preload("Questions.Options").Preload("Questions.Tags").Preload("Questions").Find(&exam)

	// 试卷信息 + 用户答卷 => 批改后试卷
	var examResult response.ExamResult

	examResult.ID = uint(examAnswerId)
	examResult.Title = exam.Title
	examResult.TotalScore = 0
	examResult.Score = 0
	examResult.StartTime = examAnswer.StartTime
	examResult.EndTime = examAnswer.EndTime

	for _, question := range exam.Questions {

		isRight := true
		var options []model.QuestionOptionResult

		var examAnswerQuestion model.ExamAnswerQuestions

		global.DB.Where(&model.ExamAnswerQuestions{
			ExamAnswerId: examAnswer.ID,
			QuestionId:   question.ID,
		}).Preload("QuestionOptions").First(&examAnswerQuestion)

		for _, option := range question.Options {
			isChoosed := false

			for _, userOption := range examAnswerQuestion.QuestionOptions {
				if option.ID == userOption.QuestionOptionId {
					isChoosed = true
				}
			}

			if (isChoosed && !option.IsRight) || (!isChoosed && option.IsRight) {
				isRight = false
			}

			options = append(options, model.QuestionOptionResult{
				Title: option.Title,
				// IsRight:   option.IsRight,
				IsChoosed: isChoosed,
			})
		}

		examResult.TotalScore++
		if isRight {
			examResult.Score++
		} else {
			// TODO: write in wrongbook
		}

		examResult.Questions = append(examResult.Questions, model.QuestionResult{
			// Title: question.Title,
			// Analysis: question.Analysis,
			// Type: question.Type,
			IsRight: isRight,
			// Options: options,
		})
	}

	// fmt.Println(examAnswer.ID, examResult.TotalScore, examResult.Score)
	return [2]uint{examResult.TotalScore, examResult.Score}, nil
}

func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库

	file, _ := os.Open("exam_answers.csv")
	// newFile, err := os.Open("exam_answers_new.csv")
	// newFile := os.Stdout
	newFile, err := os.Create("users.csv")
	if err != nil {
		log.Println(err)
		return
	}

	r := csv.NewReader(file)
	w := csv.NewWriter(newFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(record[0])

		result, err := GetExamResult(record[0])
		if err != nil {
			continue
		}
		total := strconv.Itoa(int(result[0]))
		score := strconv.Itoa(int(result[1]))
		newRecord := append(record, total, score)
		if err := w.Write(newRecord); err != nil {
			log.Println("error writing record to csv:", err)
		}
		fmt.Println(result)
		w.Flush()
		if err := w.Error(); err != nil {
			log.Println(err)
		}
	}
}
