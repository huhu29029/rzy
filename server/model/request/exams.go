package request

type ExamReq struct {
	Title       string `json:"title"`
	QuestionIDs []uint `json:"questionIDs"`
}
