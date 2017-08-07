package app

import "github.com/sirupsen/logrus"

type str string

type status  struct {
	Ok                str

	ErrInternal       string // Generic error definition, 10000-10999
	ErrInMaintain     string
	ErrRequestFormat  string
	ErrApiDeprecated  string // Api is deprectated, need to update app
	ErrSessionExpired string
	ErrInBlacklist    string
	ErrTempBlock      string // 暂时被禁止访问
}

const STATUS = status{
	Ok: "00000",
	ErrInternal:       "10000",
	ErrInMaintain:     "10001",
	ErrRequestFormat:  "10002",
	ErrApiDeprecated:  "10003",
	ErrSessionExpired: "10004",
	ErrInBlacklist:    "10005",
	ErrTempBlock:      "10006",
}

var statusMsg = map[string]string{
	STATUS.ErrInternal:       "发生未知错误，请稍后再试～",
	STATUS.ErrInMaintain:     "服务器在维护中，请稍后再试～",
	STATUS.ErrRequestFormat:  "网络请求错误",
	STATUS.ErrApiDeprecated:  "当前版本过低，需要升级咯～",
	STATUS.ErrSessionExpired: "用户已从其他设备登录",
	STATUS.ErrInBlacklist:    "当前用户已无法访问服务",
	STATUS.ErrTempBlock:      "您已暂时被禁止访问，请稍后再试",
}

func (this *status)Msg(code string) string {
	if _v, ok := statusMsg[code]; ok {
		return _v
	}

	logrus.WithFields(logrus.Fields{
		"code":code,
	}).Error("missing code")
	return ""
}
