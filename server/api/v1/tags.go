package v1

import (
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/response"

	"github.com/gin-gonic/gin"
)

type GetChaptersResponse struct {
	Chapters []model.Chapters `json:"chapters"`
}

// @Tags Chapters
// @Summary 获取所有章节
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} GetChaptersResponse "response.Response{data: GetChaptersResponse}"
// @Router /chapters [get]
func GetChapters(c *gin.Context) {
	var chapters []model.Chapters

	global.DB.Find(&chapters)

	response.OkWithData(GetChaptersResponse{
		Chapters: chapters,
	}, c)
}
