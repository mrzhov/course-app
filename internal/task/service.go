package task

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) Create(task *Task) error {
	return s.repo.Create(task)
}

func (s *Service) GetList() ([]Task, error) {
	return s.repo.GetList()
}

func (s *Service) GetById(id uint) (Task, error) {
	return s.repo.GetById(id)
}

func (s *Service) Patch(task *Task) error {
	return s.repo.Patch(task)
}

func (s *Service) Delete(task *Task) error {
	return s.repo.Delete(task)
}
