package outist

type TUserMan struct {
	users map[string]TUser
}

func CreateUserMan() *TUserMan {
	var result = &TUserMan{}
	return result
}
