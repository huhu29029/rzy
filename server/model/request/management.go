package request

type AddQuestionOption struct {
	Title   string `json:"title"`
	IsRight bool   `json:"isRight"`
}
type AddQuestionRequest struct {
	Title     string              `json:"title"`
	Analysis  string              `json:"analysis"`
	Type      uint                `json:"type"`
	Tags      []uint              `json:"tags"`
	ChapterId uint                `json:"chapterId"`
	Options   []AddQuestionOption `json:"options"`
}

type EditQuestionOption struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	IsRight    bool   `json:"isRight"`
	QuestionId uint   `json:"questionId"`
}
type EditQuestionRequest struct {
	Title     string               `json:"title"`
	Analysis  string               `json:"analysis"`
	Type      uint                 `json:"type"`
	ChapterId uint                 `json:"chapterId"`
	Options   []EditQuestionOption `json:"options"`
	Tags      []uint               `json:"tags"`
}
