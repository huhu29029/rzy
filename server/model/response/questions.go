package response

import "online-exam/model"

type GetQuestionHistoryResponse struct {
	Total           uint                    `json:"total"`
	QuestionHistory []model.QuestionHistory `json:"questionHistory"`
}

type ManagementUsersResp struct {
	Total uint            `json:"total"`
	Users []model.SysUser `json:"users"`
}

type QuestionResultResponse struct {
	Result model.QuestionResult `json:"result"`
}

type RandomQuestionResponse struct {
	Question model.Questions `json:"question"`
}

type GetQuestions2NiepanResponse struct {
	Total     uint              `json:"total"`
	Questions []model.Questions `json:"questions"`
}

type WrongQuestionResult struct {
	model.QuestionResult
	QuestionID int `json:"questionID"`
}

type GetWrongQuestionRequest struct {
	QuestionID int `json:"id"`
	UserID     int `json:"id"`
}

type GetWrongQuestionResponse struct {
	Result []WrongQuestionResult `json:"result"`
}

type AddWrongQuestionRequest struct {
	question model.QuestionResult
}

type AddWrongQuestionResponse struct {
	QuestionID int `json:"questionID"`
}

type UpdateWrongQuestionRequest struct {
	QuestionID int `json:"questionID"`
	model.QuestionResult
}

