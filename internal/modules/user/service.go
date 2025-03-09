package user

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) Create(user *User) error {
	return s.repo.Create(user)
}

func (s *Service) GetList(users *[]User) error {
	return s.repo.GetList(users)
}

func (s *Service) GetById(user *User, id uint) error {
	return s.repo.GetById(user, id)
}

func (s *Service) Patch(user *User) error {
	return s.repo.Patch(user)
}

func (s *Service) Delete(user *User) error {
	return s.repo.Delete(user)
}
