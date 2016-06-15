package outist

import (
	"encoding/json"
	"net/http"
)

type TUserLoginRequest struct {
	// For request
	UserName string
	// For request
	UserPassword string
	// For response
	SessionID string
}

func (this *TUserLoginRequest) LoadFromHttpRequest(request *http.Request) {
	this.UserName = request.FormValue("userName")
	this.UserPassword = request.FormValue("userPassword")
}

func (this *TUserLoginRequest) ToJsonData() []byte {
	var data, result = json.Marshal(this)
	if nil != result {
		GlobalLog.Write("Error: could not marshal; result.Error = " + result.Error())
	}
	return data
}
