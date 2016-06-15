package outist

import (
	"sync"
	"time"
)

type TUserMan struct {
	usersLock    sync.RWMutex
	users        map[string]TUser
	Directory    string
	backupThread TCycleThread
}

func CreateUserMan() *TUserMan {
	var this = &TUserMan{}
	this.users = make(map[string]TUser)
	this.backupThread.Interval = time.Hour
	this.backupThread.Function = this.Backup
	return this
}

func (this *TUserMan) GetUserFileName() string {
	return "Users.json"
}

func (this *TUserMan) GetUserFilePath() string {
	return this.Directory + "/" + this.GetUserFileName()
}

func (this *TUserMan) GetUserBackupFileName() string {
	return "Users_" + FormatFileTime(time.Now()) + ".json"
}

func (this *TUserMan) GetUserBackupFilePath() string {
	return this.Directory + "/" + this.GetUserBackupFileName()
}

func (this *TUserMan) Prepare() {
	this.Load()
	if false == this.getUserByName(AdminUserName).Valid() {
		this.AddAdminUser()
		this.Save(this.GetUserFilePath())
	}
	this.backupThread.Start()
}

func (this *TUserMan) AddAdminUser() {
	var addUserResult = this.AddUser(this.CreateAdminUser())
	GlobalLog.Write("Attempted to add admin user; result = " + BoolToStr(addUserResult))
}

func (this *TUserMan) CreateAdminUser() TUser {
	var password = ReadStringFromFile(this.Directory + "/admin-password.txt")
	var admin TUser
	if len(password) > 0 {
		admin.Name = AdminUserName
		admin.SetPassword(password)
	} else {
		GlobalLog.Write("Error: empty password for admin user")
	}
	return admin
}

func (this *TUserMan) SleepSecurity() {
	time.Sleep(time.Second / 2)
}

// Locks usersLock
func (this *TUserMan) AddUser(user TUser) bool {
	var result = false
	this.usersLock.Lock()
	defer this.usersLock.Unlock()
	this.SleepSecurity()
	var alreadyExists = CheckUserMapContains(this.users, user.Name)
	if alreadyExists || user.Name == "" {
		result = false
	} else {
		this.users[user.Name] = user
		result = true
	}
	return result
}

// Locks usersLock.
// Returns SessionID. Returns empty string if failed.
func (this *TUserMan) LogIn(user TUser) string {
	var result = ""
	this.usersLock.Lock()
	defer this.usersLock.Unlock()
	this.SleepSecurity()
	var theUser = this.getUserByName(user.Name)
	var matched = theUser.Valid() && theUser.PasswordHash == user.PasswordHash
	if matched {
		theUser.SessionID = GenerateUserSessionId()
		theUser.LoginTime = time.Now()
		this.setUser(theUser)
		result = theUser.SessionID
	}
	return result
}

func (this *TUserMan) Save(filePath string) {
	GlobalLog.Write("Now saving users to file '" + filePath + "'")
	var users = this.GetUserArray()
	var text = UsersToJson(users)
	var writeResult = WriteStringToFile(filePath, text)
	if writeResult == nil {
		GlobalLog.Write("Users were saved successfully")
	} else {
		GlobalLog.Write("Error: couls not save users; writeResult = " + writeResult.Error())
	}
}

func (this *TUserMan) Load() error {
	var result error
	var filePath = this.GetUserFilePath()
	GlobalLog.Write("Now loading users from '" + filePath + "'")
	var text = ReadStringFromFile(filePath)
	if len(text) > 0 {
		var userArray = JsonToUsers(text)
		var userMap = UserArrayToMap(userArray)
		this.SetUserMap(userMap)
		result = nil
	} else {
		GlobalLog.Write("Warning: empty text loaded from file '" + filePath + "'")
	}
	return result
}

func (this *TUserMan) GetUserArray() []TUser {
	this.usersLock.RLock()
	defer this.usersLock.RUnlock()
	return UserMapToUserArray(this.users)
}

func (this *TUserMan) SetUserMap(users map[string]TUser) {
	this.usersLock.Lock()
	defer this.usersLock.Unlock()
	this.users = users
}

// Does not lock.
func (this *TUserMan) getUserByName(name string) TUser {
	var result TUser
	var value, valueFound = this.users[name]
	if valueFound {
		if value.Name == name {
			result = value
		} else {
			GlobalLog.Write("Error: value.Name != name where value.Name='" + value.Name + "' name='" + name + "'")
		}
	}
	return result
}

// Locks usersLock.
func (this *TUserMan) GetUserByName(name string) TUser {
	this.usersLock.RLock()
	defer this.usersLock.RUnlock()
	return this.getUserByName(name)
}

// Does not lock
func (this *TUserMan) setUser(user TUser) bool {
	var result = false
	var oldUser, found = this.users[user.Name]
	if found {
		if oldUser.Name == user.Name {
			this.users[user.Name] = user
		} else {
			GlobalLog.Write("Error: oldUser.Name != user.Name where oldUser.Name='" + oldUser.Name + "' user.Name='" + user.Name + "'")
		}
	}
	return result
}

// Locks usersLock.
func (this *TUserMan) CheckSession(user TUser) TUser {
	var result TUser
	this.usersLock.Lock()
	defer this.usersLock.Unlock()
	var storedUser = this.getUserByName(user.Name)
	if storedUser.Valid() {
		var sessionValid = len(storedUser.SessionID) > 0 && storedUser.SessionID == user.SessionID
		if sessionValid {
			result = storedUser
			storedUser.LastCheckTime = time.Now()
			this.users[storedUser.Name] = storedUser
		}
	}
	return result
}

func (this *TUserMan) Stop() {
	this.backupThread.Stop()
	this.Save(this.GetUserFilePath())
}

func (this *TUserMan) Backup() {
	this.Save(this.GetUserBackupFilePath())
	this.Save(this.GetUserFilePath())
}
