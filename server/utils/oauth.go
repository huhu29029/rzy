package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"online-exam/global"
	"strconv"
	"time"

	"github.com/kirinlabs/HttpRequest"
)

// 火花统一登录

type OursparkDetail struct {
	Name  string `json:"Name"`
	Role  int    `json:"Role"`
	Univ  int    `json:"University"`
	Stuno string `json:"Stuno"`
}

func GetOursparkOID(token string) (oid string, err error) {
	req := HttpRequest.NewRequest()
	ts := int(time.Now().Unix())
	signtext := fmt.Sprintf("ts=%v&token=%v", ts, token)
	resp, err := req.JSON().Post(global.CONFIG.Oauth.OursparkBaseUrl+"/api/oauth/validateToken", map[string]interface{}{
		"appid": global.CONFIG.Oauth.OursparkAppid,
		"token": token,
		"sign":  GetOursparkSign(signtext, global.CONFIG.Oauth.OursparkSecret),
		"ts":    strconv.Itoa(ts),
	})
	if err != nil {
		return "", err
	}
	type result struct {
		State bool   `json:"state"`
		Type  string `json:"type"`
		Info  string `json:"info"`
	}
	fmt.Println(resp.Content())
	r := result{}
	resp.Json(&r)
	if r.State && r.Type == "uid" {
		return r.Info, nil
	} else {
		return "", errors.New("Wrong conf")
	}
}

func GetOursaprkDetail(oid string) (OursparkDetail, error) {
	req := HttpRequest.NewRequest()
	ts := int(time.Now().Unix())
	signtext := fmt.Sprintf("ts=%v&uid=%v", ts, oid)
	resp, err := req.JSON().Post(global.CONFIG.Oauth.OursparkBaseUrl+"/api/oauth/getUserDetail", map[string]interface{}{
		"appid": global.CONFIG.Oauth.OursparkAppid,
		"UID":   oid,
		"sign":  GetOursparkSign(signtext, global.CONFIG.Oauth.OursparkSecret),
		"ts":    strconv.Itoa(ts),
	})
	if err != nil {
		return OursparkDetail{}, err
	}
	type result struct {
		State bool           `json:"state"`
		Info  OursparkDetail `json:"info"`
	}
	fmt.Println(resp.Content())
	r := result{}
	resp.Json(&r)
	if r.State {
		return r.Info, nil
	} else {
		return OursparkDetail{}, errors.New("Wrong conf")
	}
}

func GetOursparkSign(text string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(text))
	data := mac.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(data)
	return sign
}
