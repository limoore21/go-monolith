package service

type UserGetter interface {
	GetFirstName() string
	GetSurName() string
}
type UserService struct {
	UserModel UserGetter
}

func NewUserService(user UserGetter) *UserService {
	return &UserService{
		UserModel: user,
	}
}

func (s *UserService) GetFullName() string {
	firstname := s.UserModel.GetFirstName()
	lastName := s.UserModel.GetSurName()

	return firstname + " " + lastName
}
