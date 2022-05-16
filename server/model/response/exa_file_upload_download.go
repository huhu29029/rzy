package response

import "online-exam/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
