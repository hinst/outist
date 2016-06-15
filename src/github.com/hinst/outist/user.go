package outist

import (
	"bytes"
	"encoding/json"
	"time"
)

const AdminUserName = "admin"
const UserNameLengthLimit = 100
const UserPasswordLengthLimit = 1000
const UserSessionIdLength = 60

const (
	UserRegistrationSuccessResult = "success"
	UserCaptchaFailedResult       = "captchaFailed"
	UserNameEmptyResult           = "nameIsEmpty"
	UserNameTooLongResult         = "userNameTooLong"
	UserPasswordEmptyResult       = "passwordEmpty"
	UserPasswordTooLongResult     = "passwordTooLong"
	UserAlreadyExistsResult       = "alreadyExists"
	UserRegistrationErrorResult   = "error"
)

type TUser struct {
	Name          string
	PasswordHash  string
	SessionID     string
	LastCheckTime time.Time
	LoginTime     time.Time
}

func (this *TUser) SetPassword(password string) {
	this.PasswordHash = GetPasswordHashString(password)
}

func (this TUser) Valid() bool {
	return len(this.Name) > 0
}

func UserMapToUserArray(users map[string]TUser) []TUser {
	var userArray = make([]TUser, len(users))
	var index = 0
	for _, value := range users {
		userArray[index] = value
		index++
	}
	return userArray
}

func UserArrayToMap(users []TUser) map[string]TUser {
	var userMap = make(map[string]TUser)
	for userIndex, currentUser := range users {
		userMap[currentUser.Name] = currentUser
		Unuse(userIndex)
	}
	return userMap
}

func UsersToJson(users []TUser) string {
	var indentedData bytes.Buffer
	var data, marshalResult = json.Marshal(users)
	if marshalResult == nil {
		json.Indent(&indentedData, data, "", "\t")
	} else {
		GlobalLog.Write("Error: could marshall users: " + marshalResult.Error())
	}
	return string(indentedData.String())
}

func JsonToUsers(jsonText string) []TUser {
	var data = []byte(jsonText)
	var usersArray = make([]TUser, 0)
	var users = &usersArray
	var result = json.Unmarshal(data, users)
	if result != nil {
		GlobalLog.Write("Could not unmarshall users: " + result.Error())
	}
	return *users
}

func CheckUserMapContains(users map[string]TUser, userName string) bool {
	var _, result = users[userName]
	return result
}

func GenerateUserSessionId() string {
	var random = CreateRandom()
	var sessionIdText bytes.Buffer
	for i := 0; i < UserSessionIdLength; i++ {
		var number = random.Intn(10)
		var text = IntToStr(number)
		sessionIdText.Write([]byte(text))
	}
	return sessionIdText.String()
}
