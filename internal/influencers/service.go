package influencers

type Repository interface {
	GetAll() ([]Influencer, error)
	GetByID(id string) (Influencer, error)
	Create(name, lastname, email string) error
}

type Service struct {
	Repo Repository
}

func NewService(r Repository) *Service {
	return &Service{Repo: r}
}

func (s *Service) GetAll() ([]Influencer, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetByID(id string) (Influencer, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) Create(name, lastname, email string) error {
	return s.Repo.Create(name, lastname, email)
}
