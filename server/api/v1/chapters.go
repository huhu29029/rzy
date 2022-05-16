package v1

import (
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/response"

	"github.com/gin-gonic/gin"
)

type GetTagsResponse struct {
	Tags []model.Tags `json:"tags"`
}

// @Tags Tag
// @Summary 获取所有标签
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} GetTagsResponse "response.Response{data: GetTagsResponse}"
// @Router /tags [get]
func GetTags(c *gin.Context) {
	var tags []model.Tags

	global.DB.Find(&tags)

	response.OkWithData(GetTagsResponse{
		Tags: tags,
	}, c)
}
