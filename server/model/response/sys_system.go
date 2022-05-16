package response

import "online-exam/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
