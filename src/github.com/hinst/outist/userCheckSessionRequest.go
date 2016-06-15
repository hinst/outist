package outist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type TUserCheckSessionRequest struct {
	UserName  string
	SessionID string
}

func (this *TUserCheckSessionRequest) LoadFromHttpRequest(request *http.Request) bool {
	var result = false
	var requestData, readBodyResult = ioutil.ReadAll(request.Body)
	if readBodyResult == nil {
		var decodeResult = json.Unmarshal(requestData, this)
		if decodeResult == nil {
			result = true
		}
	}
	return result
}

type TUserCheckSessionResponse struct {
	UserName  string
	LoggedIn  bool
	LoginTime time.Time
}
