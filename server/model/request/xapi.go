package request

import "online-exam/model"

type XAPIData struct {
	Verb   string        `json:"verb"`
	Object model.XObject `json:"object"`
}
