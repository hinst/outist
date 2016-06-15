package outist

import "time"

type TUser struct {
	Name          string
	PasswordHash  string
	SessionIDs    string
	LastCheckTime time.Time
	LoginTime     time.Time
}

func (this *TUser) SetPassword(password string) {
	this.PasswordHash = GetPasswordHashString(password)
}
