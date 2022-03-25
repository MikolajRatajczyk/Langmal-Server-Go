package service

type LoginServiceInterface interface {
	Login(username string, password string) bool
}

func NewLoginService() LoginServiceInterface {
	return &loginService{
		hardcodedUsername: "username",
		hardcodedPassword: "password",
	}
}

//	TODO: handle more than one hardcoded user
type loginService struct {
	hardcodedUsername string
	hardcodedPassword string
}

func (ls *loginService) Login(username string, password string) bool {
	okUsername := username == ls.hardcodedUsername
	okPassword := password == ls.hardcodedPassword
	return okUsername && okPassword
}
