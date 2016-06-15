package outist

import "encoding/json"

type TUserRegistrationResponse struct {
	Name            string
	OperationResult string
}

func (this *TUserRegistrationResponse) ToJsonData() []byte {
	var data, conversionResult = json.Marshal(this)
	if conversionResult != nil {
		GlobalLog.Write("Error: could not convert to json: " + conversionResult.Error())
	}
	return data
}
