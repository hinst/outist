package outist

import "net/http"

type TUserRegistrationRequest struct {
	CaptchaId     string
	CaptchaAnswer string
	UserName      string
	UserPassword  string
}

func (this *TUserRegistrationRequest) LoadFromHttpRequest(request *http.Request) {
	this.CaptchaId = request.FormValue("captchaId")
	this.CaptchaAnswer = request.FormValue("captcha")
	this.UserName = request.FormValue("userName")
	this.UserPassword = request.FormValue("userPassword")
}
