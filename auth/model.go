package auth

type ServiceLoginRet struct {
	Code      int    `json:"code"`
	Desc      string `json:"desc"`
	Qs        string `json:"qs"`
	Sid       string `json:"sid"`
	Sign      string `json:"_sign"`
	Callback  string `json:"callback"`
	PassToken string `json:"passToken"`
	UserId    int    `json:"userId"`
	Location  string `json:"location"`
	Nonce     int64  `json:"nonce"`
	Ssecurity string `json:"ssecurity"`
}
